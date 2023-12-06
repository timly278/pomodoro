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
	Data    any    `json:"data"`
}

func ErrorResponse(err error) gin.H {
	if strings.Contains(err.Error(), "\n") {
		errs := strings.SplitAfter(err.Error(), "\n")
		maping := make(map[string]any)
		for i, e := range errs {
			maping[fmt.Sprintf("error_%d", i+1)] = e
		}
		return maping
	} else {
		return gin.H{"error": err.Error()}
	}
}

type UserResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" `
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type NewTokensResponse struct {
	RefreshToken string `json:"refresh_token"`
	RTExpireIn   int64  `json:"rt_expire_in"`
	AccessToken  string `json:"access_token"`
	ATExpireIn   int64  `json:"at_expire_in"`
}
type UserLoginResponse struct {
	NewTokens NewTokensResponse `json:"new_tokens"`
	User      UserResponse      `json:"user"`
}

type UserSettingResponse struct {
	Username    string `json:"username"`
	AlarmSound  string `json:"alarm_sound"`
	RepeatAlarm int32  `json:"repeat_alarm"`
}

func NewUserResponse(user *db.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func LoginSuccessfully(user *db.User, newToken *NewTokensResponse) *Response {
	userLogin := UserLoginResponse{
		NewTokens: *newToken,
		User:      NewUserResponse(user),
	}
	return &Response{
		Message: "login successfully",
		Data:    userLogin,
	}
}
