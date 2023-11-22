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
func getObjectId(ctx *gin.Context) (int64, error) {
	id := ctx.Param("id")
	x, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid id of object, id should be a number greater then zero")
	}
	return int64(x), nil
}
