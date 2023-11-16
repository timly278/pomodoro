package api

import (
	"fmt"
	db "pomodoro/db/sqlc"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
	Data    any    `json:"any"`
}

func errorResponse(err error) gin.H {
	errs := strings.SplitAfter(err.Error(), "\n")
	maping := make(map[string]any)
	for i, e := range errs {
		maping[fmt.Sprintf("error_%d", i+1)] = e
	}
	return maping
}

type userResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" `
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
