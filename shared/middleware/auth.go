package middleware

import (
	"errors"
	"fmt"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/shared/response"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	AUTHORIZATION_HEADER_KEY        = "authorization"
	AUTHORIZATION_TYPE_BEARER       = "bearer"
	AUTHORIZATION_PAYLOAD_KEY       = "authorization_payload"
	AUTHORIZATION_ACCESSTOKEN_KEY   = "authorization_accesstoken"
	AUTHORIZATION_USERID_KEY        = "userid"
	BLACKLIST_CONTAINS_ACCESS_TOKEN = "true"
)

// EnsureLoggedIn requires user authenticated
func EnsureLoggedIn(tokenMaker security.TokenMaker, store db.Store, redisdb *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessToken, err := parseAccessToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}
		if IsAccessTokenInBlackList(redisdb, accessToken, ctx) {
			err = errors.New("user has logged out, must login again")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		payload, err := tokenMaker.VerifyToken(accessToken, security.SUBJECT_CLAIM_ACCESS_TOKEN)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		id, err := strconv.Atoi(payload.RegisteredClaims.ID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		user, err := store.GetUserById(ctx, int64(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(err))
		}

		if user.IsBlocked {
			// add accessToken to blackList of redis
			expireAt := payload.ExpiresAt.Time
			err := redisdb.Set(ctx, accessToken, BLACKLIST_CONTAINS_ACCESS_TOKEN, time.Until(expireAt)).Err()
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(err))
			}
			err = errors.New("access token has been revoked")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		}

		ctx.Set(AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Set(AUTHORIZATION_ACCESSTOKEN_KEY, accessToken)
		ctx.Set(AUTHORIZATION_USERID_KEY, int64(id))
		ctx.Next()
	}
}

// getTokenInfo checks if the request is logged in or not
func parseAccessToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader(AUTHORIZATION_HEADER_KEY)

	if len(authHeader) == 0 {
		err := errors.New("authorization header is not provide")
		return "", err
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return "", err
	}

	authType := strings.ToLower(fields[0])
	if authType != AUTHORIZATION_TYPE_BEARER {
		err := fmt.Errorf("unsupported authorization type %s", authType)
		return "", err
	}

	accessToken := fields[1]

	return accessToken, nil
}

func IsAccessTokenInBlackList(redisdb *redis.Client, accessToken string, ctx *gin.Context) bool {

	redis, err := redisdb.Get(ctx, accessToken).Result()
	fmt.Println("redis check blacklist:", redis)

	if err != nil {
		return false
	}
	if redis == BLACKLIST_CONTAINS_ACCESS_TOKEN {
		return true
	}
	return false
}
