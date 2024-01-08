package mdw

import (
	"fmt"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// [GIN] 2024/01/04 - 13:22:29 | 200 |  111.59ms |  ::1 | POST  "/api/v1/auth/login"
// ts, http.statusCode, Latency, clientID, method, path, payload, message/error
// func Formatter(param gin.LogFormatterParams) string {

// }

// how

func (m *Middleware) Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		if raw != "" {
			path = path + "?" + raw
		}
		msg := fmt.Sprintf("[GIN] %d | %15s | %s | %s | %s |",
			ctx.Writer.Status(),
			ctx.ClientIP(),
			ctx.Request.UserAgent(),
			ctx.Request.Method,
			path,
		)
		m.logger.Info(
			msg,
			zap.Any("response", ctx.Value(response.RESPONSE_CONTEXT_KEYWORD)),
		)
	}
}
