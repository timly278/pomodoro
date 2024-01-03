package auth

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Refresh access token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			token 	body		delivery.RefreshTokenRequest true "refresh access token"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	gin.H	"Bad Request"
//	@Failure		401		{object}	gin.H "Refresh token is unauthorized"
//	@Failure		500		{object}	gin.H "Internal serever error"
//	@Router			/auth/refresh-token [post]
func (t *authHandlers) RefreshToken(ctx *gin.Context) {
	var req delivery.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	newTokens, err := t.authService.RefreshTokens(ctx, req)
	if err != nil {
		t.logger.Info("refresh token is unauthorized", zap.String("detail", err.Error()))
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "refresh token successfully",
		Data:    newTokens,
	})
}
