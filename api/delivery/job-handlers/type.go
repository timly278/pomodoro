package pomodo

import (
	"database/sql"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

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

	typeId, err := delivery.GetNumericObjectParam(ctx, "id")
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
