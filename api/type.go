package api

import (
	"database/sql"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

type createTypeRequest struct {
	Name              string `json:"name" binding:"required,alphanum"`
	Color             string `json:"color" binding:"required"`
	Goalperday        int32  `json:"goal_per_day" binding:"required,min=1"`
	Duration          int32  `json:"duration" binding:"required,min=1"`
	Shortbreak        int32  `json:"shortbreak" binding:"required,min=1"`
	Longbreak         int32  `json:"longbreak" binding:"required,min=1"`
	Longbreakinterval int32  `json:"longbreakinterval" binding:"required,min=1"`
	AutostartBreak    bool   `json:"autostart_break" binding:"required,boolean"`
}

func (server *Server) CreateNewPomoType(ctx *gin.Context) {
	var typeRequest createTypeRequest
	err := ctx.ShouldBindJSON(&typeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomotype, err := server.store.CreateNewType(ctx, db.CreateNewTypeParams{
		UserID:            getUserId(ctx),
		Name:              typeRequest.Name,
		Color:             typeRequest.Color,
		Goalperday:        typeRequest.Goalperday,
		Duration:          typeRequest.Duration,
		Shortbreak:        typeRequest.Shortbreak,
		Longbreak:         typeRequest.Longbreak,
		Longbreakinterval: typeRequest.Longbreakinterval,
		AutostartBreak:    typeRequest.AutostartBreak,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Message: "create new pomodoro type successfully",
		Data:    pomotype,
	})

}

func (server *Server) ListPomoType(ctx *gin.Context) {
	types, err := server.store.ListTypes(ctx, getUserId(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, types)
}

func (server *Server) UpdateType(ctx *gin.Context) {
	var typeRequest createTypeRequest
	err := ctx.ShouldBindJSON(&typeRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	typeId, err := getNumericObjectParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomotype, err := server.store.UpdateTypeById(ctx, db.UpdateTypeByIdParams{
		ID:                typeId,
		UserID:            getUserId(ctx),
		Name:              typeRequest.Name,
		Color:             typeRequest.Color,
		Goalperday:        typeRequest.Goalperday,
		Duration:          typeRequest.Duration,
		Shortbreak:        typeRequest.Shortbreak,
		Longbreak:         typeRequest.Longbreak,
		Longbreakinterval: typeRequest.Longbreakinterval,
		AutostartBreak:    typeRequest.AutostartBreak,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, pomotype)

}
