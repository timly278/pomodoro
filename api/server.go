package api

import (
	"pomodoro/auth"
	db "pomodoro/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	tokenMaker  auth.TokenMaker
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	const secretKey = "tulb123456"
	tokenMaker := auth.NewJwtTokenMaker(secretKey)
	return &Server{
		store:  store,
		tokenMaker: tokenMaker,
		router: gin.Default(),
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Setup() {
	router := gin.Default()

	router.POST("/register", server.CreateUser)
	router.POST("/login", server.UserLogin)

	router.POST("/test", server.Dosomething)
	router.POST("/logout", server.UserLogout)

	server.router = router
}
