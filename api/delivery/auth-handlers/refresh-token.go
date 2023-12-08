package auth

import (
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (t *authHandlers) RefreshToken(ctx *gin.Context) {
	var req delivery.RefreshTokenRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	newTokens, err := t.authService.RefreshTokens(ctx, req)
	if err != nil {
		// TODO: improve error handling
		fmt.Printf("handler-err : %v\n", err)
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "refresh token successfully",
		Data:    newTokens,
	})
}
