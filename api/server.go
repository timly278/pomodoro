package api

import (
	"fmt"
	"pomodoro/auth"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/middleware"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	tokenMaker auth.TokenMaker
	config     util.Config
	router     *gin.Engine
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := auth.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		router:     gin.Default(),
	}, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Setup() {
	router := gin.Default()

	router.POST("/register", middleware.EnsureNotLoggedIn(server.tokenMaker), server.CreateUser)
	router.POST("/login", middleware.EnsureNotLoggedIn(server.tokenMaker), server.UserLogin)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.EnsureLoggedIn(server.tokenMaker))
	authRoutes.POST("/logout", server.UserLogout)
	authRoutes.POST("/dosomething/:num", server.Dosomething)

	authRoutes.PUT("/api/users", server.UpdateUserSetting)
	authRoutes.POST("/api/types", server.CreateNewPomoType)
	authRoutes.GET("/api/types", server.ListPomoType)
	authRoutes.PUT("/api/types/:id", server.UpdateType)

	authRoutes.POST("/api/pomodoros", server.CreateNewPomodoro)
	//TODO: implement for reporting total days accessed and minutes from begining
	authRoutes.GET("/api/report", server.SimpleStatisticNumber)
	authRoutes.GET("/api/report/month", server.ListPomoByMonth)
	authRoutes.GET("/api/report/date", server.ListPomoByDate)

	//TODO: Task implementation
	server.router = router
}
