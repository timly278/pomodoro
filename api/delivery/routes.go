package delivery

import (
	mdw "pomodoro/shared/middleware"

	"github.com/gin-gonic/gin"
)

/*
	TODO: How can fx show these kind of errors????

main.Fx

	/Users/timly/Workspace/Code/Go-Project/pomodoro/main.go:56

main.main

	/Users/timly/Workspace/Code/Go-Project/pomodoro/main.go:28

runtime.main

	/usr/local/go/src/runtime/proc.go:267

Failed: missing dependencies for function "pomodoro/api/delivery".MapAuthRoutes

	/Users/timly/Workspace/Code/Go-Project/pomodoro/api/delivery/routes.go:10:

missing types:
  - *gin.Engine (did you mean to Provide it?)
  - delivery.AuthHandlers (did you mean to use *auth.authHandlers?)

[Fx] ERROR              Failed to start: missing dependencies for function "pomodoro/api/delivery".MapAuthRoutes

	/Users/timly/Workspace/Code/Go-Project/pomodoro/api/delivery/routes.go:10:

missing types:
  - *gin.Engine (did you mean to Provide it?)
  - delivery.AuthHandlers (did you mean to use *auth.authHandlers?)

2023/12/17 15:33:53 missing dependencies for function "pomodoro/api/delivery".MapAuthRoutes (/Users/timly/Workspace/Code/Go-Project/pomodoro/api/delivery/routes.go:10): missing types: *gin.Engine; delivery.AuthHandlers (did you mean *auth.authHandlers?)
*/
func MapAuthRoutes(
	route *gin.Engine,
	h AuthHandlers,
	mware *mdw.Middleware,
) {
	route.GET("/", h.Home)
	group := route.Group("api/v1/auth")
	group.POST("/refresh-token", mware.Logger(), h.RefreshToken)
	group.POST("/send-emailverification", mware.Logger(), h.SendEmailVerification)
	group.POST("/verify-code", mware.Logger(), h.VerifyCode)
	group.POST("/register", mware.Logger(), h.Register)
	group.POST("/login", mware.Logger(), h.Login)

	// TODO: not implemented feature
	group.POST("/logout", mware.EnsureLoggedIn(), h.Logout)
	group.PUT("/update-password", mware.EnsureLoggedIn(), h.UpdatePassword)

	// TODO: group.PUT("/reset-password", h.UpdatePassword) // forget password
}

func MapJobsRoutes(
	router *gin.Engine,
	h JobHandlers,
	mware *mdw.Middleware,
) {
	group := router.Group("api/v1/jobs")
	gin.Logger()
	group.Use(mware.EnsureLoggedIn())

	group.PUT("/update-user-setting", h.UpdateUserSetting) // need middleware
	group.POST("/pomodoros", h.CreateNewPomodoro)
	group.GET("/pomodoros", h.ListPomodorosByDates)

	group.POST("/types", h.CreateNewPomoType)
	group.PUT("/types/:id", h.UpdatePomoType)
	group.GET("/types", h.GetPomoType)

	group.GET("/focused-minutes", h.GetMinutesFocused)
	group.GET("/accessed-days", h.GetDaysAccessed)
}
