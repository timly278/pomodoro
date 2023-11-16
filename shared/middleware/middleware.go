package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"pomodoro/auth"
	"pomodoro/shared/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION_HEADER_KEY  = "authorization"
	AUTHORIZATION_TYPE_BEARER = "bearer"
	AUTHORIZATION_PAYLOAD_KEY = "authorization_payload"
)

// EnsureLoggedIn requires user authenticated
func EnsureLoggedIn(tokenMaker auth.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)

		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provide")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != AUTHORIZATION_TYPE_BEARER {
			err := fmt.Errorf("unsupported authorization type %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
			return

		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
			return
		}

		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}

// EnsureNotLoggedIn require user unauthenticated
func EnsureNotLoggedIn() {

}
