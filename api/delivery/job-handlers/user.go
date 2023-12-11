package jobs

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (u *jobHandlers) UpdateUserSetting(ctx *gin.Context) {
	var req delivery.UpdateUserSettingRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	newSetting, err := u.userService.UpdateUserSetting(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "update fields successfully",
		Data:    newSetting,
	})
}