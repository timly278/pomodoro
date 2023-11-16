package api

import (
	"net/http"
	"pomodoro/auth"
	"pomodoro/shared/middleware"
	"pomodoro/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) Dosomething(ctx *gin.Context) {
	num := ctx.Param("num")
	x, err := strconv.Atoi(num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	x = x * 1000000

	payload := ctx.MustGet(middleware.AUTHORIZATION_PAYLOAD_KEY).(*auth.Payload)

	ctx.JSON(http.StatusOK, gin.H{
		"x":        strconv.Itoa(x),
		"username": payload.Issuer,
	})
}
