package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"
	"strconv"
	"time"

	gomail "gopkg.in/mail.v2"

	"github.com/gin-gonic/gin"
)

const (
	VERIFICATION_CODE_LIFETIME = 4 * time.Minute
)

type verificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type verificationResponse struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

func (server *Server) EmailVerification(ctx *gin.Context) {
	var verifyRequest verificationRequest
	err := ctx.ShouldBindJSON(&verifyRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	if !server.checkCodeAndEmail(ctx, verifyRequest.Email, verifyRequest.Code) {
		err = errors.New("unaccepted code")
		ctx.JSON(http.StatusNotAcceptable, response.ErrorResponse(err))
		return
	}

	user, statusCode, err := server.emailVerificationDBUpdate(ctx, verifyRequest.Email)
	if err != nil {
		ctx.JSON(statusCode, response.ErrorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(strconv.FormatInt(user.ID, 10), server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}
	res := verificationResponse{
		Token:  accessToken,
		UserID: user.ID,
		Email:  verifyRequest.Email,
	}

	// TODO: redirect to logged-in page
	ctx.JSON(http.StatusOK, response.Response{
		Message: "login successfully",
		Data:    res,
	})
}

func (server *Server) emailVerificationDBUpdate(ctx *gin.Context, email string) (user db.User, statusCode int, err error) {
	user, err = server.store.UpdateVerifyEmail(ctx, db.UpdateVerifyEmailParams{
		Email:         email,
		EmailVerified: true,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			statusCode = http.StatusNotFound
			return
		}
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusAccepted
	return
}

func (server *Server) checkCodeAndEmail(ctx *gin.Context, email, code string) bool {

	redisEmail, err := server.redisdb.Get(ctx, code).Result()
	if err != nil {
		return false
	}
	if redisEmail != email {
		return false
	}
	return true
}

// sendEmailVerification should be invoked in a new goroutine
func (server *Server) sendEmailVerification(ctx *gin.Context, email, password string) {
	// TODO: improve error logging for the sake of error investigation
	code := util.RandomVerificationCode()

	err := server.redisdb.Set(ctx, code, email, VERIFICATION_CODE_LIFETIME).Err()
	if err != nil {
		log.Printf("err: %v\nemail: %s, password: %s\n", email, password, err)
		return
	}

	body := formEmailBody(email, password, code)

	err = server.EmailSendingExecution(email, body)
	if err != nil {
		log.Printf("err: %v\nemail: %s, password: %s\n", email, password, err)
	}
}

// TODO: how to render html code on email?
func formEmailBody(email, password string, code string) string {
	return fmt.Sprintf(`<html>
    <body>
        <h2>Your verified code: %s</h2>
        <br>
        <p>Your Email Account: %s</p>
        <p>Your Password: %s</p>
    </body>
	</html>`, code, email, password)
}

func (server *Server) isEmailExisted(ctx *gin.Context, email string) (user db.User, statusCode int, err error) {

	user, err = server.store.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			statusCode = http.StatusNotFound
		}
		statusCode = http.StatusInternalServerError
	}
	statusCode = http.StatusFound
	return
}

func (server *Server) EmailSendingExecution(userEmail, body string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", server.config.AppEmail)

	// Set E-Mail receivers
	m.SetHeader("To", userEmail)

	// Set E-Mail subject
	m.SetHeader("Subject", "Verify email CODE of Pomodoro.com")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// send E-Mail
	if err := server.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
