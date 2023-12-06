package pomodo

import (
	"net/http"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type yearMonthRequest struct {
	Year  int32 `form:"year" binding:"required"`
	Month int32 `form:"month" binding:"required,min=1,max=12"`
}

type statisticNumberResponse struct {
	DaysAccessed   int64 `json:"days_accessed"`
	Minutesfocused int64 `json:"minutes_focused"`
}

// GET("/api/report")
func (pomo *pomoHandlers) SimpleStatisticNumber(ctx *gin.Context) {
	var timeRequest yearMonthRequest
	err := ctx.ShouldBindQuery(&timeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userID := delivery.GetUserId(ctx)

	days, err := pomo.getDaysAccessedInMonth(ctx, userID, timeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	minutes, err := pomo.getMinutesFocusedInMonth(ctx, userID, timeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, statisticNumberResponse{
		DaysAccessed:   days,
		Minutesfocused: minutes,
	})

}

func (pomo *pomoHandlers) getDaysAccessedInMonth(ctx *gin.Context, userID int64, timeRequest yearMonthRequest) (int64, error) {
	daysAccessed, err := pomo.store.GetDaysAccessedInMonth(ctx, db.GetDaysAccessedInMonthParams{
		UserID:  userID,
		MonthID: timeRequest.Month,
		Year:    timeRequest.Year,
	})
	if err != nil {
		return 0, err
	}

	return daysAccessed, nil
}

func (pomo *pomoHandlers) getMinutesFocusedInMonth(ctx *gin.Context, userID int64, timeRequest yearMonthRequest) (int64, error) {
	minutes, err := pomo.store.GetMinutesFocusedInMonth(ctx, db.GetMinutesFocusedInMonthParams{
		UserID:  userID,
		MonthID: timeRequest.Month,
		Year:    timeRequest.Year,
	})
	if err != nil {
		if strings.Contains(err.Error(), "NULL to int64") {
			return 0, nil
		}
		return 0, err
	}

	return minutes, nil
}
