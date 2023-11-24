package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"time"

	"github.com/gin-gonic/gin"
)

type createPomodoroRequest struct {
	TypeID      int64 `json:"type_id" binding:"required,min=1"`
	TaskID      int64 `json:"task_id"`
	FocusDegree int32 `json:"focus_degree" binding:"required,min=1,max=5"`
}

func (server *Server) CreateNewPomodoro(ctx *gin.Context) {
	var pomoRequest createPomodoroRequest

	err := ctx.ShouldBindJSON(&pomoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomodoro, err := server.createPomodoro(ctx, pomoRequest)
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

func (server *Server) createPomodoro(ctx *gin.Context, pomoRequest createPomodoroRequest) (pomodoro db.Pomodoro, err error) {
	var taskID sql.NullInt64
	userId := getUserId(ctx)
	if pomoRequest.TaskID == 0 {
		pomodoro, err = server.store.CreatePomodoroWithNoTask(ctx, db.CreatePomodoroWithNoTaskParams{
			UserID:      userId,
			TypeID:      pomoRequest.TypeID,
			FocusDegree: pomoRequest.FocusDegree,
		})
	} else {
		taskID.Scan(pomoRequest.TaskID)
		pomodoro, err = server.store.CreatePomodoroWithTask(ctx, db.CreatePomodoroWithTaskParams{
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
func (server *Server) ListPomoByMonth(ctx *gin.Context) {

	var timeRequest yearMonthRequest
	err := ctx.ShouldBindQuery(&timeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	date := time.Date(int(timeRequest.Year), time.Month(timeRequest.Month), 0, 0, 0, 0, 0, time.Local)
	numberOfDate := date.Day()
	rsp := make([][]db.GetPomodoroByDateRow, numberOfDate)

	userID := getUserId(ctx)

	// time.Date might get less efficient than forming a date string like: '2023-11-23'
	for i := 0; i < numberOfDate; i++ {
		date = date.AddDate(0, 0, 1)
		params := db.GetPomodoroByDateParams{
			UserID:    userID,
			Limit:     30, //TODO: what if there are more than 30 pomodoros a day?
			Offset:    0,
			QueryDate: date,
		}
		pomo, err := server.store.GetPomodoroByDate(ctx, params)
		if err != nil {
			if err == sql.ErrNoRows {
				// scan another day of the month
				continue
			}
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
			return
		}
		rsp[i] = append(rsp[i], pomo...)
	}

	ctx.JSON(http.StatusOK, rsp)

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

func (server *Server) ListPomoByDate(ctx *gin.Context) {

	var pomoRequest listPomoByDateRequest
	err := ctx.ShouldBindQuery(&pomoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userID := getUserId(ctx)
	dbQueryParams := db.GetPomodoroByDateParams{
		UserID:    userID,
		Limit:     pomoRequest.PageSize,
		Offset:    (pomoRequest.PageID - 1) * pomoRequest.PageSize,
		QueryDate: pomoRequest.DateTime,
	}

	pomos, err := server.store.GetPomodoroByDate(ctx, dbQueryParams)
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
