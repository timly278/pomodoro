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

	user, err := u.authService.CreateUser(ctx, req)
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

func (u *authHandlers) UpdateUserSetting(ctx *gin.Context) {
	var req delivery.UpdateUserSettingRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	newSetting, err := u.authService.UpdateUserSetting(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "update setting successfully",
		Data:    newSetting,
	})
}
