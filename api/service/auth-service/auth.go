package authservice

import (
	"pomodoro/api/service"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type authService struct {
	userService service.UserService
	store       db.Store
	tokenMaker  security.TokenMaker
	dialer      *gomail.Dialer
	redisdb     *redis.Client
	config      *util.Config
}

func NewAuthService(
	store db.Store,
	userService service.UserService,
	tokenMaker security.TokenMaker,
	redisdb *redis.Client,
	dialer *gomail.Dialer,
	conf *util.Config) service.AuthService {

	return &authService{
		userService: userService,
		store:       store,
		tokenMaker:  tokenMaker,
		dialer:      dialer,
		redisdb:     redisdb,
		config:      conf}
}
