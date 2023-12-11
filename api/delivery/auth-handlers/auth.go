package auth

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	authservice "pomodoro/api/service/auth-service"
	userservice "pomodoro/api/service/user-service"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type authHandlers struct {
	authService service.AuthService
	userService service.UserService
}

var _ delivery.AuthHandlers = (*authHandlers)(nil)

func NewAuthHandlers(store db.Store,
	tokenMaker security.TokenMaker,
	redisdb *redis.Client,
	dialer *gomail.Dialer,
	conf *util.Config) *authHandlers {

	authService := authservice.NewAuthService(store, tokenMaker, redisdb, dialer, conf)
	userService := userservice.NewUserService(store)
	return &authHandlers{authService: authService, userService: userService}
}
