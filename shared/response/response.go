package response

import (
	db "pomodoro/db/sqlc"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
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

func NewUserResponse(user *db.User) *UserResponse {
	return &UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func Response(ctx *gin.Context, mess string, data any) response {
	ctx.Set(RESPONSE_CONTEXT_KEYWORD, mess)
	return response{Message: mess, Data: data}
}
