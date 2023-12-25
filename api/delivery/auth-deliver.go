package delivery

import "github.com/gin-gonic/gin"

type AuthHandlers interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	SendEmailVerification(ctx *gin.Context)
	VerifyCode(ctx *gin.Context)
}
