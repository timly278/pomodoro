package api

import (
	db "pomodoro/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	return &Server{
		store:  store,
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
