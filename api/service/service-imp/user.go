package logging

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
	"pomodoro/util"

	"github.com/redis/go-redis/v9"
)

type UserLogging struct {
	store      db.Store
	redisdb    *redis.Client // to create black list of logging out
	tokenMaker TokenLogging
}

func NewUserLogging(store db.Store, redisdb *redis.Client) *UserLogging {
	return &UserLogging{store: store, redisdb: redisdb}
}

func (u *UserLogging) CreateUser(ctx context.Context, req delivery.CreateUserRequest) (*db.User, error) {
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

func (u *UserLogging) Login(ctx context.Context, req delivery.LoginRequest) (tokens *response.NewTokensResponse, code int, err error) {
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

	tokens, err = u.tokenMaker.newTokens(ctx, user.ID)
	if err != nil {
		return
	}
	return
}

// logout
func (u *UserLogging) Logout(ctx context.Context, accessToken string) {

}

// update password

func (u *UserLogging) UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error) {
	user, err := u.store.UpdateUserSetting(ctx, db.UpdateUserSettingParams{
		ID:          userId,
		Username:    req.Username,
		AlarmSound:  req.AlarmSound,
		RepeatAlarm: req.RepeatAlarm,
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

func (u *UserLogging) GetUserByMail(ctx context.Context, mail string) (user *db.User, code int, err error) {
	*user, err = u.store.GetUserByEmail(ctx, mail)
	if err != nil {
		if err == sql.ErrNoRows {
			code = http.StatusNotFound
		}
		code = http.StatusInternalServerError
	}
	code = http.StatusFound
	return
}
