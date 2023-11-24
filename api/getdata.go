package api

import (
	"fmt"
	"net/http"
	"pomodoro/auth"
	"pomodoro/shared/middleware"
	"pomodoro/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Dosomething just for testing
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

// getObjectId returns error if the request is bad
func getNumericObjectParam(ctx *gin.Context, key string) (int64, error) {
	id := ctx.Param(key)
	x, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid key, %s should be a number greater than zero", key)
	}
	return int64(x), nil
}

func getUserId(ctx *gin.Context) int64 {
	return ctx.MustGet(middleware.AUTHORIZATION_USERID_KEY).(int64)
}
