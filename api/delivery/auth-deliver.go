package delivery

import "github.com/gin-gonic/gin"

type EmailHandlers interface {
	SendCode(ctx *gin.Context)
	Verify(ctx *gin.Context)
}

type UserHandlers interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}

type TokenHandlers interface {
	RefreshToken(ctx *gin.Context)
}

type AuthHandlers interface {
	EmailHandlers
	UserHandlers
	TokenHandlers
}