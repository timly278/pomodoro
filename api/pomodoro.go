package api

import (
	"database/sql"
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
		ctx.JSON(http.StatusBadRequest, response.ErrorMultiResponse(err))
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
	var rsp []db.GetPomodoroByDateRow

	monthID, err := getObjectId(ctx)
	if err != nil || (monthID > 12 || monthID < 1) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userID := getUserId(ctx)
	present := time.Now()


	for i := range time.
	params := db.GetPomodoroByDateParams{
		UserID: userID,
		Limit: 30,
		Offset: 0,
		QueryDate: time.Date(present.Year(), time.Month(monthID), 1, 0,0,0,0,present.Location()),
	}
	
	pomo, err := server.store.GetPomodoroByDate(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {

		}
	}

}

// GET("/api/report/:date")
func (server *Server) ListPomoByDate(ctx *gin.Context) {

}
