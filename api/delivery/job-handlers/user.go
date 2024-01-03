package jobs

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

// GetPomoType godoc
//
//	@Summary		Update User Setting
//	@Description	Update user setting
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			userSetting	body		delivery.UpdateUserSettingRequest	true	"Update user setting"
//	@Success		200			{object}	response.Response	"Update user setting successfully"
//	@Failure		400			{object}	gin.H				"Bad Request"
//	@Failure		500			{object}	gin.H 				"Internal serever error"
//	@Router			/jobs/update-user-setting [put]
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
