package api

import (
	"crypto/tls"
	"fmt"
	"pomodoro/auth"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/middleware"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
)

type Server struct {
	store      db.Store
	tokenMaker auth.TokenMaker
	config     util.Config
	router     *gin.Engine
	dialer     *gomail.Dialer
	redisdb    *redis.Client
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := auth.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Create Redis Client
	redisDb := redis.NewClient(&redis.Options{
		Addr:     config.RedisClientAddress,
		Password: config.RedisDbPassword,
		DB:       config.RedisDb,
	})

	// Settings for SMTP server
	dialer := gomail.NewDialer(config.AppSmtpHost, config.AppSmtpPort, config.AppEmail, config.AppPassword)
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		router:     gin.Default(),
		dialer:     dialer,
		redisdb:    redisDb,
	}, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Setup() {
	router := gin.Default()

	router.POST("/create-user", server.CreateUser)
	router.POST("/login", middleware.EnsureNotLoggedIn(server.tokenMaker), Login)
	
	router.POST("/send-email", server.SendCode)
	router.GET("/verify-email", server.Verify) // and login

	router.POST("/refresh-token", server.RefreshToken)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.EnsureLoggedIn(server.tokenMaker))

	authRoutes.POST("/logout", server.UserLogout)

	api := authRoutes.Group("/api")
	api.PUT("/users", server.UpdateUserSetting)
	api.POST("/pomodoros", server.CreateNewPomodoro)

	types := api.Group("/types")
	types.POST("/", server.CreateNewPomoType)
	types.GET("/", server.ListPomoType)
	types.PUT("/:id", server.UpdateType)

	report := api.Group("/report")
	report.GET("/", server.SimpleStatisticNumber)
	report.GET("/month", server.ListPomoByMonth)
	report.GET("/date", server.ListPomoByDate)
	//TODO: implement for reporting total days accessed and minutes from begining

	//TODO: Task implementation
	server.router = router
}
