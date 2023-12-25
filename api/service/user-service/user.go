package userservice

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"
	"strings"
)

type userService struct {
	store db.Store
}

func NewUserService(store db.Store) service.UserService {
	return &userService{store: store}
}

func (u *userService) CreateUser(ctx context.Context, req *delivery.CreateUserRequest) (*db.User, int, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	createUserParams := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}
	user, err := u.store.CreateUser(ctx, createUserParams)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return nil, http.StatusConflict, errors.New("email existed")
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func (u *userService) GetUserByMail(ctx context.Context, mail string) (*db.User, int, error) {

	user, err := u.store.GetUserByEmail(ctx, mail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusFound, err
}

// update password
// TODO: improve this function to be able to reveive any change from clients
func (u *userService) UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error) {
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

func (u *userService) GetUserById(ctx context.Context, userId int64) (*db.User, int, error) {
	user, err := u.store.GetUserById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusFound, err
}
