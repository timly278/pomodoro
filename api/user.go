package api

import (
	"database/sql"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Email    string `json:"email" binding:"required,email"`
}

// CreateUser - user signing up
func (server *Server) CreateUser(ctx *gin.Context) {
	var newUser createUserRequest
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	createUserParams := db.CreateUserParams{
		Username:       newUser.Username,
		HashedPassword: hashedPassword,
		Email:          newUser.Email,
	}
	createdUser, err := server.store.CreateUser(ctx, createUserParams)
	if err != nil {
		//todo: handle detail error of pq here
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	rsp := response.NewUserResponse(createdUser)
	ctx.JSON(http.StatusOK, response.Response{
		Message: "register successfully",
		Data:    rsp,
	})
}

type userLoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
}

type userLoginResponse struct {
	AccessToken string                `json:"access_token"`
	User        response.UserResponse `json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
	var userLogin userLoginRequest

	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, userLogin.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	err = util.VerifyPassword(userLogin.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, 5*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	res := userLoginResponse{
		AccessToken: accessToken,
		User:        response.NewUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "login successfully",
		Data:    res,
	})

}

func (server *Server) UserLogout(ctx *gin.Context) {

}

