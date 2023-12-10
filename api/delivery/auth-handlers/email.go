package auth

import (
	"net/http"
	"pomodoro/api/delivery"
	_ "pomodoro/docs"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
)

// SendCode godoc
//
//	@Summary		Send email for verification code
//	@Description	Send verification code
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			email body  delivery.SendCodeRequest  true "lybatu send code"
//	@Success		200			{object}	response.Response
//	@Failure		400			{string}	string	"Bad Request"
//	@Failure		406			{string}	string	"Unacceptable"
//	@Router			/api/v1/auth/send-email [post]
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

// Verify godoc
//
//	@Summary		Verify code
//	@Description	Verify code that sent over email
//	@Tags			lybatuTags
//	@Accept			json
//	@Produce		json
//	@Param			verificationRequest	body delivery.VerificationRequest true "verify code"
//	@Success		200	{object}		response.Response "email has been verified successfully"
//	@Failure		400	{string}		"Bad Request"
//	@Router			/api/v1/auth [get]
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
