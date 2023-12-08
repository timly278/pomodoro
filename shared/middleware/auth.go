package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"pomodoro/security"
	"pomodoro/shared/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION_HEADER_KEY  = "authorization"
	AUTHORIZATION_TYPE_BEARER = "bearer"
	AUTHORIZATION_PAYLOAD_KEY = "authorization_payload"
	AUTHORIZATION_USERID_KEY  = "userid"
)

// EnsureLoggedIn requires user authenticated
func EnsureLoggedIn(tokenMaker security.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		payload, err := getPayload(tokenMaker, ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		id, err := strconv.Atoi(payload.RegisteredClaims.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		// user, code, err := authS.GetUserById(ctx, int64(id))
		// if code != http.StatusFound {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		// }
		// _ = user
		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Set(AUTHORIZATION_USERID_KEY, int64(id))
		ctx.Next()
	}
}

// EnsureNotLoggedIn require user unauthenticated
func EnsureNotLoggedIn(tokenMaker security.TokenMaker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, err := getPayload(tokenMaker, ctx)
		if err == nil {
			// user already logged in
			// redirect to another page then abort the request
			ctx.Abort()
		}

		ctx.Next()
	}

}

// isLoggedIn checks if the request is logged in or not
func getPayload(tokenMaker security.TokenMaker, ctx *gin.Context) (*security.Payload, error) {
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
	payload, err := tokenMaker.VerifyToken(accessToken, security.SUBJECT_CLAIM_ACCESS_TOKEN)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
