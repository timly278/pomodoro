package delivery

import (
	"github.com/gin-gonic/gin"
)


func MapAuthRoutes(group *gin.RouterGroup, h AuthHandlers) {
	group.POST("/refresh-token", h.RefreshToken)
	group.POST("/send-email", h.SendCode)
	group.GET("/verify-email", h.Verify)
	group.POST("/create-user", h.CreateUser)
	group.POST("/login", h.Login)
	group.POST("/logout", h.Logout)
	group.PUT("/update-password", h.UpdatePassword)
	// TODO: group.PUT("/reset-password", h.UpdatePassword)
	group.PUT("/update-usersetting", h.UpdateUserSetting)
}

func  MapPomoRoutes(group *gin.RouterGroup, h PomoHandlers) {

}