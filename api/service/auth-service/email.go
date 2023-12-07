package authservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "pomodoro/db/sqlc"
	"pomodoro/util"
	"time"

	gomail "gopkg.in/mail.v2"
)

const (
	VERIFICATION_CODE_LIFETIME = 4 * time.Minute
)

func (ev *authService) Send(ctx context.Context, userEmail string) error {
	_, err := ev.redisdb.Get(ctx, userEmail).Result()
	if err == nil {
		err = errors.New("email spam, verification code has created and sent")
		return err
	}

	code := util.RandomVerificationCode()

	go ev.send(ctx, userEmail, code)

	return nil
}

func (ev *authService) Verify(ctx context.Context, email, code string) (bool, error) {
	if !ev.compareCode(ctx, email, code) {
		err := errors.New("unaccepted code")
		return false, err
	}

	err := ev.saveDatabase(ctx, email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ev *authService) saveDatabase(ctx context.Context, email string) error {

	_, err := ev.store.UpdateUser(ctx, db.UpdateUserParams{
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
		EmailVerified: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	return err
}

func (ev *authService) compareCode(ctx context.Context, email, code string) bool {

	redisCode, err := ev.redisdb.Get(ctx, email).Result()
	if err != nil {
		return false
	}
	if code != redisCode {
		return false
	}
	return true
}

func formEmailBody(code string) string {
	return fmt.Sprintf(`<html>
    <body>
        <h2>Your verified code: %s</h2>
        <br>
    </body>
	</html>`, code)
}
func (ev *authService) send(ctx context.Context, userEmail, code string) {
	m := gomail.NewMessage()
	body := formEmailBody(code)

	// Set E-Mail sender
	m.SetHeader("From", ev.config.AppEmail)

	// Set E-Mail receivers
	m.SetHeader("To", userEmail)

	// Set E-Mail subject
	m.SetHeader("Subject", "Verify email CODE of Pomodoro.com")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", body)

	// send E-Mail
	if err := ev.dialer.DialAndSend(m); err != nil {
		// TODO: log error
		return
	}

	err := ev.redisdb.Set(ctx, userEmail, code, VERIFICATION_CODE_LIFETIME).Err()
	if err != nil {
		// TODO: log error
		return
	}

}
