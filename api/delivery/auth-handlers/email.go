package auth

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

func (eh *authHandlers) SendCode(ctx *gin.Context) {

	// TODO: do I need to write some middleware to protect this kind of API?

	var req delivery.SendCodeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	err = eh.authService.Send(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "verification email has sent",
		Data:    req.Email,
	})

}

func (eh *authHandlers) Verify(ctx *gin.Context) {
	var req delivery.VerificationRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ok, err := eh.authService.Verify(ctx, req.Email, req.Code)
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "email has been verified successfully",
		Data:    req.Email,
	})
}
