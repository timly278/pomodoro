package delivery

import (
	"pomodoro/security"
	"pomodoro/shared/middleware"

	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(group *gin.RouterGroup, h AuthHandlers, tokenMaker security.TokenMaker) {
	group.POST("/refresh-token", h.RefreshToken)
	group.POST("/send-email", h.SendCode)
	group.GET("/verify-email", h.Verify)
	group.POST("/create-user", h.CreateUser)
	group.POST("/login", h.Login)

	group.POST("/logout", middleware.EnsureLoggedIn(tokenMaker), h.Logout) // need middleware
	group.PUT("/update-password", middleware.EnsureLoggedIn(tokenMaker), h.UpdatePassword)
	// TODO: group.PUT("/reset-password", h.UpdatePassword) // forget password
}

func MapJobsRoutes(group *gin.RouterGroup, h JobHandlers, tokenMaker security.TokenMaker) {
	group.Use(middleware.EnsureLoggedIn(tokenMaker))
	
	group.PUT("/update-user-setting", h.UpdateUserSetting) // need middleware
	group.POST("/pomodoros", h.CreateNewPomodoro)
	group.GET("/pomodoros", h.ListPomodorosByDates)

	group.POST("/types", h.CreateNewPomoType)
	group.PUT("/types/:id", h.UpdatePomoType)
	group.GET("/types", h.GetPomoType)

	group.GET("/focused-minutes", h.GetMinutesFocused)
	group.GET("/accessed-days", h.GetDaysAccessed)
}
