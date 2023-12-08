package delivery

import (
	"fmt"
	"pomodoro/shared/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) int64 {
	return ctx.MustGet(middleware.AUTHORIZATION_USERID_KEY).(int64)
}

// getObjectId returns error if the request is bad
func GetNumericObjectParam(ctx *gin.Context, key string) (int64, error) {
	id := ctx.Param(key)
	x, err := strconv.Atoi(id)
	if err != nil || x <= 0 {
		return 0, fmt.Errorf("invalid key, %s should be a number and greater than zero", key)
	}
	return int64(x), nil
}
