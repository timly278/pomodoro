package server

import (
	"fmt"
	"pomodoro/api/delivery"
	"pomodoro/api/delivery/auth-handlers"
	pomodo "pomodoro/api/delivery/job-handlers"
	db "pomodoro/db/sqlc"
	_ "pomodoro/docs"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"

	// gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger embed files
	swaggerFiles "github.com/swaggo/files"
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

	auths := router.Group("api/v1/auth")
	delivery.MapAuthRoutes(auths, authHandlers, s.tokenMaker)

	jobs := router.Group("api/v1/jobs")
	delivery.MapJobsRoutes(jobs, jobHandlers, s.tokenMaker)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(address)
}
