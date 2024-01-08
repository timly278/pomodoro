package delivery

import (
	mdw "pomodoro/shared/middleware"

	"github.com/gin-gonic/gin"
)

func GetUserId(ctx *gin.Context) int64 {
	return ctx.MustGet(mdw.AUTHORIZATION_USERID_KEY).(int64)
}
