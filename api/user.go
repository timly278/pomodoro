package api

import (
	"database/sql"
	"net/http"
	db "pomodoro/db/sqlc"
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(createdUser)
	ctx.JSON(http.StatusOK, response{
		Message: "register successfully",
		Data:    rsp,
	})
}

type userLoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
}

type userLoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
	var userLogin userLoginRequest

	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, userLogin.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.VerifyPassword(userLogin.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, 1*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	res := userLoginResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response{
		Message: "login successfully",
		Data:    res,
	})

}

func (server *Server) UserLogout(ctx *gin.Context) {

}

func (server *Server) Dosomething(ctx *gin.Context) {

}
