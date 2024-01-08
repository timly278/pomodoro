package auth

import (
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (u *authHandlers) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Home": "Hello World!",
	})
}

// Register godoc
//
//	@Summary		New user registers
//	@Description	Create new user and send verification code to email
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			NewUser 	body		delivery.CreateUserRequest true "Create new user"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	gin.H	"Bad Request"
//	@Failure		406			{object}	gin.H "email spam, verification code has created and sent"
//	@Failure		409			{object}	gin.H "email existed"
//	@Failure		500			{object}	gin.H "internal serever error"
//	@Router			/auth/register [post]
func (u *authHandlers) Register(ctx *gin.Context) {
	var req delivery.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	user, statusCode, err := u.userService.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(statusCode, response.ErrorResponse(ctx, err))
		return
	}

	err = u.authService.SendEmailVerification(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(ctx, err))
		return
	}

	rsp := response.NewUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		fmt.Sprintf("%s | %s | register successfully, waiting for email verification.", req.Email, req.Username),
		rsp,
	))
}

// Login godoc
//
//	@Summary		Loggin user
//	@Description	Loggin user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user 	body		delivery.LoginRequest true "user login"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	gin.H	"Bad Request"
//	@Failure		403		{object}	gin.H "password does not match"
//	@Failure		406		{object}	gin.H "email has not verified"
//	@Failure		500		{object}	gin.H "internal serever error"
//	@Router			/auth/login [post]
func (u *authHandlers) Login(ctx *gin.Context) {
	var req delivery.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	rsp, code, err := u.authService.Login(ctx, &req)
	if err != nil {
		ctx.JSON(code, response.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		fmt.Sprintf("%s | loggin successfully", req.Email),
		&rsp,
	))
}

func (u *authHandlers) Logout(ctx *gin.Context) {
	err := u.authService.Logout(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response(ctx, "loggout successfully", ""))
}

func (u *authHandlers) UpdatePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, response.Response(
		ctx,
		"Not implemented: the feature UpdatePassword",
		"",
	))
}
