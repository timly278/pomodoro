package auth

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

// RefreshToken godoc
//
//	@Summary		Refresh access token
//	@Description	Refresh access token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			token 	body		delivery.RefreshTokenRequest true "refresh access token"
//	@Success		200		{object}	response.response
//	@Failure		400		{object}	gin.H	"Bad Request"
//	@Failure		401		{object}	gin.H "Refresh token is unauthorized"
//	@Failure		500		{object}	gin.H "Internal serever error"
//	@Router			/auth/refresh-token [post]
func (t *authHandlers) RefreshToken(ctx *gin.Context) {
	var req delivery.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}
	newTokens, err := t.authService.RefreshTokens(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		"refresh token successfully",
		newTokens,
	))
}
