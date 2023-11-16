package response

import (
	"fmt"
	db "pomodoro/db/sqlc"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"any"`
}

func ErrorResponse(err error) gin.H {
	errs := strings.SplitAfter(err.Error(), "\n")
	maping := make(map[string]any)
	for i, e := range errs {
		maping[fmt.Sprintf("error_%d", i+1)] = e
	}
	return maping
}

type UserResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" `
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
