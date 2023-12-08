package authservice

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"
)

func (u *authService) CreateUser(ctx context.Context, req *delivery.CreateUserRequest) (*db.User, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	createUserParams := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}
	user, err := u.store.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *authService) Login(ctx context.Context, req *delivery.LoginRequest) (*response.NewTokensResponse, int, error) {
	user, code, err := u.GetUserByMail(ctx, req.Email)
	if code != http.StatusFound {
		return nil, code, err
	}

	err = util.VerifyPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, http.StatusForbidden, errors.New("password does not match")
	}

	if !user.EmailVerified {
		return nil, http.StatusNotAcceptable, errors.New("email has not verified")
	}

	tokens, err := u.newTokens(ctx, user.ID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return tokens, http.StatusOK, nil
}

// logout
func (u *authService) Logout(ctx context.Context, accessToken string) {

}

// update password

// TODO: move this method to pomo-service package????
// TODO: improve this function to be able to reveive any change from clients
func (u *authService) UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error) {
	user, err := u.store.UpdateUser(ctx, db.UpdateUserParams{
		ID: sql.NullInt64{
			Int64: userId,
			Valid: true,
		},
		Username: sql.NullString{
			String: req.Username,
			Valid:  true,
		},
		AlarmSound: sql.NullString{
			String: req.AlarmSound,
			Valid:  true,
		},
		RepeatAlarm: sql.NullInt32{
			Int32: req.RepeatAlarm,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	rsp := response.UserSettingResponse{
		Username:    user.Username,
		AlarmSound:  user.AlarmSound,
		RepeatAlarm: user.RepeatAlarm,
	}

	return &rsp, nil
}

func (u *authService) GetUserByMail(ctx context.Context, mail string) (*db.User, int, error) {

	user, err := u.store.GetUserByEmail(ctx, mail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusFound, err
}

func (u *authService) GetUserById(ctx context.Context, userId int64) (*db.User, int, error) {
	user, err := u.store.GetUserById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusFound, err
}
