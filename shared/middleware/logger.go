package mdw

import (
	"github.com/gin-gonic/gin"
)

// [GIN] 2024/01/04 - 13:22:29 | 200 |  111.59ms |  ::1 | POST  "/api/v1/auth/login"
// ts, http.statusCode, Latency, clientID, method, path, payload, message/error
// func Formatter(param gin.LogFormatterParams) string {

// }

// how

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// TODO: WHat if we log also on registers APIs like refresh-token
		// ther is no payload to log?

		// if raw != "" {
		// 	path = path + "?" + raw
		// }
		// msg := fmt.Sprintf("[GIN] %d | %15s | %s | %s",
		// 	ctx.Writer.Status(),
		// 	ctx.ClientIP(),
		// 	ctx.Request.Method,
		// 	path,
		// )
		// m.logger.Info(
		// 	msg,
		// 	zap.Any("payload", payload),
		// )
	}
}
