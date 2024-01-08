package jobs

import (
	"database/sql"
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateNewPomoType godoc
//
//	@Summary		Create New Pomodoro Types
//	@Description	Create new pomodoro types
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			NewPomoType 	body		delivery.CreateNewTypeRequest true "New pomodoro type"
//	@Success		200				{object}	response.Response
//	@Failure		400				{object}	gin.H	"Bad Request"
//	@Failure		500				{object}	gin.H "Internal serever error"
//	@Router			/jobs/types [post]
func (pomo *jobHandlers) CreateNewPomoType(ctx *gin.Context) {
	var req delivery.CreateNewTypeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	pomotype, err := pomo.jobService.CreateNewType(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		"create new pomodoro type successfully",
		&pomotype,
	))

}

// GetPomoType godoc
//
//	@Summary		Get Pomodoro Types
//	@Description	Get all pomodoros type of this user
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Response
//	@Failure		500			{object}	gin.H "Internal serever error"
//	@Router			/jobs/types [get]
func (pomo *jobHandlers) GetPomoType(ctx *gin.Context) {
	types, err := pomo.jobService.GetTypes(ctx, delivery.GetUserId(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		"get all types oke",
		types,
	))
}

// GetPomoType godoc
//
//	@Summary		Update Pomodoro Types
//	@Description	Update pomodoros type of this user
//	@Tags			jobs
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Type ID"
//	@Param			type	body		delivery.CreateNewTypeRequest	true	"Update pomo type"
//	@Success		200		{object}	db.Type	"Update type successfully"
//	@Failure		400		{object}	gin.H	"Bad Request"
//	@Failure		404		{object}	gin.H 	"Not found user_id or type_id"
//	@Failure		500		{object}	gin.H 	"Internal serever error"
//	@Router			/jobs/types [put]
func (pomo *jobHandlers) UpdatePomoType(ctx *gin.Context) {
	var req delivery.CreateNewTypeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	typeId, err := getNumericObjectParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}
	userId := delivery.GetUserId(ctx)
	pomotype, err := pomo.jobService.UpdateType(ctx, userId, typeId, &req)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, pomotype)

}

// getObjectId returns error if the request is bad
func getNumericObjectParam(ctx *gin.Context, key string) (int64, error) {
	id := ctx.Param(key)
	x, err := strconv.Atoi(id)
	if err != nil || x <= 0 {
		return 0, fmt.Errorf("invalid key, %s should be a number and greater than zero", key)
	}
	return int64(x), nil
}
