package service

import (
	"context"
	delivery "pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
)


type UserService interface {
	CreateUser(ctx context.Context, req *delivery.CreateUserRequest) (*db.User, int, error)
	UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error)
	GetUserByMail(ctx context.Context, mail string) (*db.User, int, error)
	GetUserById(ctx context.Context, userId int64) (user *db.User, code int, err error)
	// forget password
}

type AuthService interface {
	RefreshTokens(ctx context.Context, req delivery.RefreshTokenRequest) (*response.NewTokensResponse, error)
	Login(ctx context.Context, req *delivery.LoginRequest) (tokens *response.NewTokensResponse, code int, err error)
	// Logout
	SendEmailVerification(ctx context.Context, userEmail string) error
	VerifyCode(ctx context.Context, email, code string) (bool, error)
}

type JobsService interface {
	CreatePomodoro(ctx context.Context, userId int64, req *delivery.CreatePomodoroRequest) (*db.Pomodoro, error)
	GetPomodorosByDates(ctx context.Context, userId int64, req *delivery.GetPomodorosRequest) ([]db.GetPomodorosRow, error)
	GetMinutesFocused(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error)
	GetDaysAccessed(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error)
	GetTypes(ctx context.Context, userId int64) ([]db.Type, error)
	CreateNewType(ctx context.Context, userId int64, req *delivery.CreateNewTypeRequest) (db.Type, error)
	UpdateType(ctx context.Context, userId int64, typeId int64, req *delivery.CreateNewTypeRequest) (db.Type, error)
	//type

}
