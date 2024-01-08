package auth

import (
	"fmt"
	"net/http"
	"pomodoro/api/delivery"
	_ "pomodoro/docs"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

// SendEmailVerification godoc
//
//	@Summary		Send email for verification code
//	@Description	Send verification code
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			email	body		delivery.SendCodeRequest	true	"send code"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	gin.H	"Bad Request"
//	@Failure		406		{object}	gin.H 	"email spam, verification code has created and sent"
//	@Failure		500		{object}	gin.H 	"internal serever error"
//	@Router			/auth/send-emailverification [post]
func (eh *authHandlers) SendEmailVerification(ctx *gin.Context) {

	// TODO: do I need to write some middleware to protect this kind of API?

	var req delivery.SendCodeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	err = eh.authService.SendEmailVerification(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		fmt.Sprintf("verification email has sent to %s ", req.Email),
		req.Email,
	))

}

// VerifyCode godoc
//
//	@Summary		Verify email verification code
//	@Description	Verify code that sent over email
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			Email&Code	body		delivery.VerificationRequest	true	"verify code"
//	@Failure		400			{object}	gin.H				"Bad Request"
//	@Success		200			{object}	response.Response	"email has been verified successfully"
//	@Router			/auth/verify-code [post]
func (eh *authHandlers) VerifyCode(ctx *gin.Context) {
	var req delivery.VerificationRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	ok, err := eh.authService.VerifyCode(ctx, req.Email, req.Code)
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(ctx, err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		ctx,
		fmt.Sprintf("%s has been verified successfully", req.Email),
		req.Email,
	))
}
