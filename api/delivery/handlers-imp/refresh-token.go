package handlers

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	logging "pomodoro/api/service/service-imp"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/shared/response"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
)

type TokenHandlers struct {
	tokenService service.TokenServer
}

var _ delivery.TokenHandlers = (*TokenHandlers)(nil)

func NewTokenHandlers(store db.Store, tokenMaker security.TokenMaker, conf *util.Config) *TokenHandlers {
	tokenService := logging.NewTokenLogging(store, tokenMaker, conf)
	return &TokenHandlers{tokenService: tokenService}
}

func (t *TokenHandlers) RefreshToken(ctx *gin.Context) {
	var req delivery.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	newTokens, err := t.tokenService.RefreshTokens(ctx, req)
	if err != nil {
		// TODO: improve error handling
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "refresh token successfully",
		Data:    newTokens,
	})
}
