package pomodo

import (
	"database/sql"
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (pomo *jobHandlers) CreateNewPomodoro(ctx *gin.Context) {
	var req delivery.CreatePomodoroRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomodoro, err := pomo.pomoService.CreatePomodoro(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, pomodoro)
}

func (pomo *jobHandlers) ListPomodorosByDates(ctx *gin.Context) {

	var req delivery.GetPomodorosRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	pomos, err := pomo.pomoService.GetPomodorosByDates(ctx, userID, &req)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: fmt.Sprintf("list pomodoros from %v to %v\n", req.FromDate, req.ToDate),
		Data:    pomos,
	})

}

func (pomo *jobHandlers) GetMinutesFocused(ctx *gin.Context) {
	var req delivery.GetStatisticRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	minutesFocused, err := pomo.pomoService.GetMinutesFocused(ctx, userID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"from_date":       req.FromDate,
		"to_date":         req.ToDate,
		"minutes_focused": minutesFocused,
	})
}

func (pomo *jobHandlers) GetDaysAccessed(ctx *gin.Context) {
	var req delivery.GetStatisticRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	daysAccessed, err := pomo.pomoService.GetDaysAccessed(ctx, userID, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"from_date":     req.FromDate,
		"to_date":       req.ToDate,
		"days_accessed": daysAccessed,
	})
}
