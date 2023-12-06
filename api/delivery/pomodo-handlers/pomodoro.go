package pomodo

import (
	"database/sql"
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (pomo *pomoHandlers) CreateNewPomodoro(ctx *gin.Context) {
	var pomoRequest delivery.CreatePomodoroRequest

	err := ctx.ShouldBindJSON(&pomoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomodoro, err := pomo.createPomodoro(ctx, pomoRequest)
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

func (pomo *pomoHandlers) createPomodoro(ctx *gin.Context, pomoRequest delivery.CreatePomodoroRequest) (pomodoro db.Pomodoro, err error) {
	var taskID sql.NullInt64
	userId := delivery.GetUserId(ctx)
	if pomoRequest.TaskID == 0 {
		pomodoro, err = pomo.store.CreatePomodoroWithNoTask(ctx, db.CreatePomodoroWithNoTaskParams{
			UserID:      userId,
			TypeID:      pomoRequest.TypeID,
			FocusDegree: pomoRequest.FocusDegree,
		})
	} else {
		taskID.Scan(pomoRequest.TaskID)
		pomodoro, err = pomo.store.CreatePomodoroWithTask(ctx, db.CreatePomodoroWithTaskParams{
			UserID:      userId,
			TypeID:      pomoRequest.TypeID,
			TaskID:      taskID,
			FocusDegree: pomoRequest.FocusDegree,
		})
	}
	return
}

// ListPomoType should receive page_id and page_size?

// GET("/api/report/month/:id")
// response the whole data of the specified month
func (pomo *pomoHandlers) ListPomoByMonth(ctx *gin.Context) {

	var timeRequest yearMonthRequest
	err := ctx.ShouldBindQuery(&timeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	rsp, err := pomo.listPomodoroByMonth(ctx, timeRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (pomo *pomoHandlers) listPomodoroByMonth(ctx *gin.Context, timeRequest yearMonthRequest) ([][]db.GetPomodoroByDateRow, error) {
	date := time.Date(int(timeRequest.Year), time.Month(timeRequest.Month), 0, 0, 0, 0, 0, time.Local)
	numberOfDate := date.Day()
	rsp := make([][]db.GetPomodoroByDateRow, numberOfDate)

	userID := delivery.GetUserId(ctx)

	// time.Date might get less efficient than forming a date string like: '2023-11-23'
	for i := 0; i < numberOfDate; i++ {
		date = date.AddDate(0, 0, 1)
		params := db.GetPomodoroByDateParams{
			UserID:    userID,
			Limit:     50,
			Offset:    0,
			QueryDate: date,
		}
		pomo, err := pomo.store.GetPomodoroByDate(ctx, params)
		if err != nil {
			if err == sql.ErrNoRows {
				// scan another day of the month
				continue
			}
			return nil, err
		}
		rsp[i] = append(rsp[i], pomo...)
	}
	return rsp, nil
}

type listPomoByDateRequest struct {
	// NOTE: for using time validation here
	// using `validate:"required,DateOnly"`
	// or `validate:"required,rfc3339"`
	// or `time_format:"2006-02-01"` is the same.

	DateTime time.Time `form:"date_time" binding:"required" validate:"required,rfc3339"`
	PageID   int32     `form:"page_id" binding:"required,min=1"`
	PageSize int32     `form:"page_size" binding:"required,min=5,max=50"`
}

func (pomo *pomoHandlers) ListPomoByDate(ctx *gin.Context) {

	var pomoRequest listPomoByDateRequest
	err := ctx.ShouldBindQuery(&pomoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userID := delivery.GetUserId(ctx)
	dbQueryParams := db.GetPomodoroByDateParams{
		UserID:    userID,
		Limit:     pomoRequest.PageSize,
		Offset:    (pomoRequest.PageID - 1) * pomoRequest.PageSize,
		QueryDate: pomoRequest.DateTime,
	}

	pomos, err := pomo.store.GetPomodoroByDate(ctx, dbQueryParams)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: fmt.Sprintf("%v", pomoRequest.DateTime),
		Data:    pomos,
	})

}
