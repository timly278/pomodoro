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
