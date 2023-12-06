package service

import (
	"context"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
)

type EmailVerifier interface {
	Send(ctx context.Context, userEmail string) error
	Verify(ctx context.Context, email, code string) (bool, error)
}

type User interface {
	CreateUser(ctx context.Context, req delivery.CreateUserRequest) (*db.User, error)
	GetUserByMail(ctx context.Context, mail string) (*db.User, int, error)
	UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error)
}

type TokenServer interface {
	RefreshTokens(ctx context.Context, req delivery.RefreshTokenRequest) (*response.NewTokensResponse, error)
}
