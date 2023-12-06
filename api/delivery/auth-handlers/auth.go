package auth

import (
	"pomodoro/api/service"
	authservice "pomodoro/api/service/auth-service"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type authHandlers struct {
	authService service.AuthService
}

func NewAuthHandlers(store db.Store,
	tokenMaker security.TokenMaker,
	redisdb *redis.Client,
	dialer *gomail.Dialer,
	conf *util.Config) *authHandlers {

	authService := authservice.NewAuthService(store, tokenMaker, redisdb, dialer, conf)
	return &authHandlers{authService: authService}
}
