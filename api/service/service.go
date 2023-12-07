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

type AuthService interface {
	EmailVerifier
	User
	TokenServer
}

type PomodoService interface {
	CreatePomodoro(ctx context.Context, userId int64, req *delivery.CreatePomodoroRequest) (*db.Pomodoro, error)
	GetPomodorosByDates(ctx context.Context, userId int64, req *delivery.GetPomodorosRequest) ([]db.GetPomodorosRow, error)
	GetMinutesFocused(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error)
	GetDaysAccessed(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error)
	GetTypes(ctx context.Context, userId int64) ([]db.Type, error)
	CreateNewType(ctx context.Context, userId int64, req *delivery.CreateNewTypeRequest) (db.Type, error)
	UpdateType(ctx context.Context, userId int64, typeId int64, req *delivery.CreateNewTypeRequest) (db.Type, error)
}
