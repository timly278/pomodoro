package delivery

import "github.com/gin-gonic/gin"

type EmailHandlers interface {
	SendEmailVerification(ctx *gin.Context)
	VerifyCode(ctx *gin.Context)
}

type UserHandlers interface {
	Register(ctx *gin.Context)
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
