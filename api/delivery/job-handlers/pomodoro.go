package jobs

import (
	"database/sql"
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

// CreateNewPomodoro godoc
//
//	@Summary		Create New Pomodoro
//	@Description	Create new pomodoro
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			NewPomo 	body		delivery.CreatePomodoroRequest true "New pomodoro"
//	@Success		200			{object}	db.Pomodoro
//	@Failure		400			{object}	gin.H	"Bad Request"
//	@Failure		404			{object}	gin.H "Not found user_id or type_id"
//	@Failure		500			{object}	gin.H "Internal serever error"
//	@Router			/jobs/pomodoros [post]
func (pomo *jobHandlers) CreateNewPomodoro(ctx *gin.Context) {
	var req delivery.CreatePomodoroRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomodoro, err := pomo.jobService.CreatePomodoro(ctx, delivery.GetUserId(ctx), &req)
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

// ListPomodorosByDates godoc
//
//	@Summary		List Pomodoros
//	@Description	Get pomodoros from date to date
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			GetPomos 	body		delivery.GetPomodorosRequest true "Get pomodoros"
//	@Success		200			{array}		db.Pomodoro
//	@Failure		400			{object}	gin.H	"Bad Request"
//	@Failure		500			{object}	gin.H "Internal serever error"
//	@Router			/jobs/pomodoros [get]
func (pomo *jobHandlers) ListPomodorosByDates(ctx *gin.Context) {

	var req delivery.GetPomodorosRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	pomos, err := pomo.jobService.GetPomodorosByDates(ctx, userID, &req)
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: fmt.Sprintf("list pomodoros from %v to %v\n", req.FromDate, req.ToDate),
		Data:    pomos,
	})

}

// GetMinutesFocused godoc
//
//	@Summary		Get Minutes Focused
//	@Description	Get total minutes user has spent on pomodoro from date to date
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			DateToDate 	body		delivery.GetStatisticRequest true "Get minutes"
//	@Success		200			{object}	gin.H 			"Successfully get minutes focused"
//	@Failure		400			{object}	gin.H	"Bad Request"
//	@Failure		500			{object}	gin.H "Internal serever error"
//	@Router			/jobs/focused-minutes [get]
func (pomo *jobHandlers) GetMinutesFocused(ctx *gin.Context) {
	var req delivery.GetStatisticRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	minutesFocused, err := pomo.jobService.GetMinutesFocused(ctx, userID, &req)
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

// GetDaysAccessed godoc
//
//	@Summary		Get Days Accessed
//	@Description	Get total days user has accessed on pomodoro from date to date
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			DateToDate 	body		delivery.GetStatisticRequest true "Get days"
//	@Success		200			{object}	gin.H 			"Successfully get days accessed"
//	@Failure		400			{object}	gin.H	"Bad Request"
//	@Failure		500			{object}	gin.H "Internal serever error"
//	@Router			/jobs/accessed-days [get]
func (pomo *jobHandlers) GetDaysAccessed(ctx *gin.Context) {
	var req delivery.GetStatisticRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	userID := delivery.GetUserId(ctx)
	daysAccessed, err := pomo.jobService.GetDaysAccessed(ctx, userID, &req)
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
