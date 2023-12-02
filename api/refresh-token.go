package api

import (
	"errors"
	"net/http"
	"pomodoro/auth"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	REFRESH_TOKEN_FAKE = "refresh token reuse detected or expired"
)

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// use JSON to exchange refresh token
func (server *Server) RefreshToken(ctx *gin.Context) {

	// get refresh token over json
	var refreshToken refreshTokenRequest

	err := ctx.ShouldBindJSON(&refreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	// verify RefreshToken
	payload, err := server.tokenMaker.VerifyToken(refreshToken.RefreshToken, auth.SUBJECT_CLAIM_REFRESH_TOKEN)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	userId, err := strconv.Atoi(payload.RegisteredClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	// check on DB
	user, err := server.store.GetUserById(ctx, int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	statusCode, err := server.checkDbRefreshToken(ctx, user, refreshToken.RefreshToken)
	if statusCode != http.StatusAccepted {
		ctx.JSON(statusCode, response.ErrorResponse(err))
		return
	}

	// issue new Access Token and Refresh Token
	newTokens, err := server.issueNewTokens(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	rsp := response.LoginSuccessfully(&user, &newTokens)
	ctx.JSON(http.StatusOK, *rsp)

}

func (server *Server) issueNewTokens(ctx *gin.Context, userId int64) (rsp response.NewTokensResponse, err error) {
	accessToken, err := server.newAccessToken(userId)
	if err != nil {
		return
	}

	refreshToken, err := server.newRefreshToken(userId)
	if err != nil {
		return
	}

	params := db.UpdateRefreshTokenParams{
		ID:           userId,
		RefreshToken: refreshToken,
	}
	_, err = server.store.UpdateRefreshToken(ctx, params)
	if err != nil {
		return
	}

	rsp = response.NewTokensResponse{
		RefreshToken: refreshToken,
		RTExpireIn:   int64(server.config.RefreshTokenDuration.Seconds()),
		AccessToken:  accessToken,
		ATExpireIn:   int64(server.config.AccessTokenDuration.Seconds()),
	}
	return
}

func (server *Server) newRefreshToken(userId int64) (refreshToken string, err error) {
	return server.tokenMaker.CreateToken(
		strconv.FormatInt(userId, 10),
		auth.SUBJECT_CLAIM_REFRESH_TOKEN,
		server.config.RefreshTokenDuration)
}
func (server *Server) newAccessToken(userId int64) (refreshToken string, err error) {
	return server.tokenMaker.CreateToken(
		strconv.FormatInt(userId, 10),
		auth.SUBJECT_CLAIM_ACCESS_TOKEN,
		server.config.AccessTokenDuration)
}

func (server *Server) checkDbRefreshToken(ctx *gin.Context, user db.User, refreshToken string) (statusCode int, err error) {

	if user.RefreshToken != REFRESH_TOKEN_FAKE {

		if user.RefreshToken != refreshToken {
			statusCode, err = server.removeRefreshTokenDB(ctx, user.ID)

		} else {
			statusCode = http.StatusAccepted
			err = nil
		}

	} else {
		// deny and ask user re-login
		statusCode = http.StatusForbidden
	}

	if statusCode == http.StatusForbidden {
		err = errors.New("refresh token reuse detected or expired, user must login again")
	}
	return
}

func (server *Server) removeRefreshTokenDB(ctx *gin.Context, userId int64) (statusCode int, err error) {
	_, err = server.store.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
		ID:           userId,
		RefreshToken: REFRESH_TOKEN_FAKE,
	})
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	statusCode = http.StatusForbidden
	return
}

// use cookie to send request refresh token and response
