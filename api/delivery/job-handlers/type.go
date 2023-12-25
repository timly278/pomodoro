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

func (pomo *jobHandlers) CreateNewPomoType(ctx *gin.Context) {
	var req delivery.CreateNewTypeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	pomotype, err := pomo.jobService.CreateNewType(ctx, delivery.GetUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Message: "create new pomodoro type successfully",
		Data:    pomotype,
	})

}

func (pomo *jobHandlers) GetPomoType(ctx *gin.Context) {
	types, err := pomo.jobService.GetTypes(ctx, delivery.GetUserId(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Message: "all types",
		Data:    types,
	})
}

func (pomo *jobHandlers) UpdatePomoType(ctx *gin.Context) {
	var req delivery.CreateNewTypeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	typeId, err := getNumericObjectParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userId := delivery.GetUserId(ctx)
	pomotype, err := pomo.jobService.UpdateType(ctx, userId, typeId, &req)
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

// getObjectId returns error if the request is bad
func getNumericObjectParam(ctx *gin.Context, key string) (int64, error) {
	id := ctx.Param(key)
	x, err := strconv.Atoi(id)
	if err != nil || x <= 0 {
		return 0, fmt.Errorf("invalid key, %s should be a number and greater than zero", key)
	}
	return int64(x), nil
}
