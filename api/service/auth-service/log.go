package authservice

import (
	"context"
	"errors"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/shared/response"
	"pomodoro/util"
)

func (u *authService) Login(ctx context.Context, req *delivery.LoginRequest) (*response.NewTokensResponse, int, error) {
	user, code, err := u.userService.GetUserByMail(ctx, req.Email)
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

// foget password
