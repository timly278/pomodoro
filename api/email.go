package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/util"
	"time"

	gomail "gopkg.in/mail.v2"

	"github.com/gin-gonic/gin"
)

const (
	VERIFICATION_CODE_LIFETIME = 4 * time.Minute
)

func (server *Server) VerifyEmail(ctx *gin.Context) {
	// get the CODE over URI

	// check the CODE on Redis along with user id
}

// TODO: improve error logging for the sake of error investigation
func (server *Server) sendEmailVerification(ctx *gin.Context, userRequest createUserRequest) {

	code := util.RandomVerificationCode()

	err := server.redisdb.Set(ctx, code, userRequest.Email, VERIFICATION_CODE_LIFETIME).Err()
	if err != nil {
		log.Println(userRequest, err)
		return
	}

	body := formEmailBody(userRequest, code)

	err = server.EmailSendingExecution(userRequest.Email, body)
	if err != nil {
		log.Println(userRequest, err)
	}
}

// TODO: how to render html code on email?
func formEmailBody(userRequest createUserRequest, code string) string {
	return fmt.Sprintf(`<html>
    <body>
        <h2>Your verified code: %s</h2>
        <br>
        <p>Your Username: %s</p>
        <p>Your Email: %s</p>
        <p>Your Password: %s</p>
    </body>
	</html>`, code, userRequest.Username, userRequest.Email, userRequest.Password)
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
