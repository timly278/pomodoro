package delivery

import "github.com/gin-gonic/gin"

type JobHandlers interface {
	CreateNewPomodoro(ctx *gin.Context)
	ListPomodorosByDates(ctx *gin.Context)
	GetMinutesFocused(ctx *gin.Context)
	GetDaysAccessed(ctx *gin.Context)
	CreateNewPomoType(ctx *gin.Context)
	GetPomoType(ctx *gin.Context)
	UpdatePomoType(ctx *gin.Context)
}
