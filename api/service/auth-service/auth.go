package authservice

import (
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type authService struct {
	store      db.Store
	tokenMaker security.TokenMaker
	dialer     *gomail.Dialer
	redisdb    *redis.Client
	config     *util.Config
}

func NewAuthService(
	store db.Store,
	tokenMaker security.TokenMaker,
	redisdb *redis.Client,
	dialer *gomail.Dialer,
	conf *util.Config) *authService {
	return &authService{
		store:      store,
		tokenMaker: tokenMaker,
		dialer:     dialer,
		redisdb:    redisdb,
		config:     conf}
}
