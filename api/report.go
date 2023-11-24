package api

import (
	"fmt"
	"net/http"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (server *Server) SimpleStatisticNumber(ctx *gin.Context) {
	monthID, err := getNumericObjectParam(ctx, "month")
	if err != nil || (monthID > 12 || monthID < 1) {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	fmt.Printf("monthID = %d\n", monthID)
}
