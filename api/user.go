package api

import (
	"errors"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" `
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

// CreateUser - user signing up
func (server *Server) CreateUser(ctx *gin.Context) {
	var newUser createUserRequest
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, statusCode, err := server.isEmailExisted(ctx, newUser.Email)
	if statusCode == http.StatusInternalServerError {
		ctx.JSON(statusCode, response.ErrorResponse(err))
		return
	} else if statusCode == http.StatusFound {
		if user.EmailVerified { //email is verified
			err = errors.New("error: account has existed")
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
			return
		}
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	var message string
	if user != (db.User{}) {
		user, statusCode, err = server.updateUserPassword(ctx, user.ID, hashedPassword)
		if err != nil {
			ctx.JSON(statusCode, response.ErrorResponse(err))
			return
		}
		message = "user existed and password has been updated successfully"
	} else {
		createUserParams := db.CreateUserParams{
			Username:       newUser.Username,
			HashedPassword: hashedPassword,
			Email:          newUser.Email,
		}
		user, err = server.store.CreateUser(ctx, createUserParams)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
			return
		}
		message = "create new user successfully"
	}

	go server.sendEmailVerification(ctx, newUser.Email, newUser.Password)

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response{
		Message: message,
		Data:    rsp,
	})
}

func (server *Server) updateUserPassword(ctx *gin.Context, userID int64, hashedPass string) (user db.User, statusCode int, err error) {

	user, err = server.store.UpdatePassword(ctx, db.UpdatePasswordParams{
		ID:                userID,
		HashedPassword:    hashedPass,
		PasswordChangedAt: time.Now(),
	})
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	return user, http.StatusOK, nil
}

type userLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=12"`
}

type userLoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

// UserLogin
func (server *Server) UserLogin(ctx *gin.Context) {
	var userLogin userLoginRequest

	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, statusCode, err := server.isEmailExisted(ctx, userLogin.Email)
	if statusCode != http.StatusFound {
		ctx.JSON(statusCode, response.ErrorResponse(err))
		return
	}

	err = util.VerifyPassword(userLogin.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	if !user.EmailVerified {
		go server.sendEmailVerification(ctx, user.Email, userLogin.Password)
		// redirect user to verification email page
		err = errors.New("waiting for user's email verification")
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(strconv.FormatInt(user.ID, 10), server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	res := userLoginResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "login successfully",
		Data:    res,
	})

}

// TODO: UserLogout
func (server *Server) UserLogout(ctx *gin.Context) {
	// create blacklist to store logged out token
	// implement blacklist on Redis to take advantage of speed

	// if you use oauth to login with google for example, you dont need to let user logout
	// todo: add login using oauth2.0
	// todo: do we really need to implement setUserStatus middleware handler?
}

type updateUserSettingRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	AlarmSound  string `json:"alarm_sound" binding:"required"`
	RepeatAlarm int32  `json:"repeat_alarm" binding:"required,alphanum,min=1,max=10"`
}

func (server *Server) UpdateUserSetting(ctx *gin.Context) {
	var settingRequest updateUserSettingRequest

	err := ctx.ShouldBindJSON(&settingRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, err := server.store.UpdateUserSetting(ctx, db.UpdateUserSettingParams{
		ID:          getUserId(ctx),
		Username:    settingRequest.Username,
		AlarmSound:  settingRequest.AlarmSound,
		RepeatAlarm: settingRequest.RepeatAlarm,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response{
		Message: "update setting successfully",
		Data:    rsp,
	})
}
