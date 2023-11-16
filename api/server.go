package api

import (
	"fmt"
	"pomodoro/auth"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.Store
	tokenMaker auth.TokenMaker
	router     *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	const secretKey = "tulb123456789tulb123456789tulb123456789tulb123456789"
	tokenMaker, err := auth.NewJwtTokenMaker(secretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	return &Server{
		store:      store,
		tokenMaker: tokenMaker,
		router:     gin.Default(),
	}, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Setup() {
	router := gin.Default()

	router.POST("/register", server.CreateUser)
	router.POST("/login", server.UserLogin)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.EnsureLoggedIn(server.tokenMaker))
	authRoutes.POST("/dosomething/:num", server.Dosomething)
	authRoutes.POST("/logout", server.UserLogout)

	server.router = router
}
