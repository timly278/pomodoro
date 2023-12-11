package delivery

import (
	"pomodoro/shared/middleware"

	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) int64 {
	return ctx.MustGet(middleware.AUTHORIZATION_USERID_KEY).(int64)
}
