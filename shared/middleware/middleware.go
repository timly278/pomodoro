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

		payload, err := isLoggedIn(tokenMaker, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}

// EnsureNotLoggedIn require user unauthenticated
func EnsureNotLoggedIn(tokenMaker auth.TokenMaker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, err := isLoggedIn(tokenMaker, ctx)
		if err == nil {
			// user already logged in
			// redirect to another page then abort the request
			ctx.Abort()
		}

		ctx.Next()
	}

}

// isLoggedIn checks if the request is logged in or not
func isLoggedIn(tokenMaker auth.TokenMaker, ctx *gin.Context) (*auth.Payload, error) {
	authHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)

	if len(authHeader) == 0 {
		err := errors.New("authorization header is not provide")
		return nil, err
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return nil, err
	}

	authType := strings.ToLower(fields[0])
	if authType != AUTHORIZATION_TYPE_BEARER {
		err := fmt.Errorf("unsupported authorization type %s", authType)
		return nil, err
	}

	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
