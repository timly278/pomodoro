package pomodo

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	jobservice "pomodoro/api/service/job-service"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

type jobHandlers struct {
	jobService service.JobsService
}

var _ delivery.JobHandlers = (*jobHandlers)(nil)

func NewPomoHandlers(store db.Store) *jobHandlers {
	job := jobservice.NewJobService(store)
	return &jobHandlers{jobService: job}
}

func (u *jobHandlers) UpdateUserSetting(ctx *gin.Context) {
	var req delivery.UpdateUserSettingRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	newSetting, err := u.jobService.UpdateUserSetting(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "update fields successfully",
		Data:    newSetting,
	})
}
