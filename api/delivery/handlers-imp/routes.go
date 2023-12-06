package handlers

import (
	"pomodoro/api/delivery"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(router *gin.Engine) *Router {
	return &Router{router: router}
}
func (r *Router) MapTokenRoutes(group *gin.RouterGroup, h delivery.TokenHandlers) {
	group.POST("/refresh-token", h.RefreshToken)
}

func (r *Router) MapMailRoutes(group *gin.RouterGroup, h delivery.EmailHandlers) {
	group.POST("/send-email", h.SendCode)
	group.GET("/verify-email", h.Verify)
}

func (r *Router) MapUserRoutes(group *gin.RouterGroup, h delivery.UserHandlers) {
	group.POST("/create-user", h.CreateUser)
	group.POST("/login", h.Login)
	group.POST("/logout", h.Logout)
	group.PUT("/update-password", h.UpdatePassword)
	// TODO: group.PUT("/reset-password", h.UpdatePassword)
	group.PUT("/update-usersetting", h.UpdateUserSetting)
}
