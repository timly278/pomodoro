package server

import (
	"fmt"
	"pomodoro/api/delivery"
	"pomodoro/api/delivery/auth-handlers"
	pomodo "pomodoro/api/delivery/job-handlers"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type Server struct {
	store      db.Store
	tokenMaker security.TokenMaker
	config     *util.Config
	dialer     *gomail.Dialer
	redisdb    *redis.Client
	// logger
}

func NewServer(store db.Store, config *util.Config, dialer *gomail.Dialer, redisDb *redis.Client) (*Server, error) {

	tokenMaker, err := security.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		dialer:     dialer,
		redisdb:    redisDb,
	}, nil
}

func (s *Server) Run(address string) {
	router := gin.Default()
	authHandlers := auth.NewAuthHandlers(s.store, s.tokenMaker, s.redisdb, s.dialer, s.config)
	jobHandlers := pomodo.NewPomoHandlers(s.store)
	delivery.MapAuthRoutes(router.Group("api/v1/auth"), authHandlers, s.tokenMaker)
	delivery.MapJobsRoutes(router.Group("api/v1/job"), jobHandlers, s.tokenMaker)

	router.Run(address)
}
