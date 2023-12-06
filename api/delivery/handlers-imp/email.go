package handlers

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	logging "pomodoro/api/service/service-imp"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type emailHandlers struct {
	mailService service.EmailVerifier
	userService service.User
}

var _ delivery.EmailHandlers = (*emailHandlers)(nil)

func NewEmailHandlers(store db.Store, redisdb *redis.Client, dialer *gomail.Dialer, from string) *emailHandlers {
	mailService := logging.NewEmailVerifier(store, redisdb, dialer, from)
	userService := logging.NewUserLogging(store, redisdb)

	return &emailHandlers{mailService: mailService, userService: userService}
}

func (eh *emailHandlers) SendCode(ctx *gin.Context) {

	// TODO: do I need to write some middleware to protect this kind of API?

	var req delivery.SendCodeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	err = eh.mailService.Send(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "verification email has sent",
		Data:    req.Email,
	})

}

func (eh *emailHandlers) Verify(ctx *gin.Context) {
	var req delivery.VerificationRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ok, err := eh.mailService.Verify(ctx, req.Email, req.Code)
	if !ok {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "email has been verified successfully",
		Data:    req.Email,
	})
}
