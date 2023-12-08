package auth

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (u *authHandlers) CreateUser(ctx *gin.Context) {
	var req delivery.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, err := u.authService.CreateUser(ctx, &req)
	if err != nil {
		// TODO: handle specific error i.e sql.NoRowErr
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	rsp := response.NewUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response{
		Message: "create new user successfully",
		Data:    rsp,
	})
}

func (u *authHandlers) Login(ctx *gin.Context) {
	var req delivery.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	rsp, code, err := u.authService.Login(ctx, &req)
	if err != nil {
		ctx.JSON(code, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "loggin successfully",
		Data:    rsp,
	})
}

func (u *authHandlers) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, response.Response{
		Message: "not implemented feature",
	})
}

func (u *authHandlers) UpdatePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, response.Response{
		Message: "not implemented feature",
	})
}


