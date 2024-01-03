package authservice

import (
	"context"
	"errors"
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/security"
	"pomodoro/shared/middleware"
	"pomodoro/shared/response"
	"pomodoro/util"
	"time"
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

// Logout service
func (u *authService) Logout(ctx context.Context) error {
	// add Access Token to blacklist
	payload := ctx.Value(middleware.AUTHORIZATION_PAYLOAD_KEY).(*security.Payload)
	accessToken := ctx.Value(middleware.AUTHORIZATION_ACCESSTOKEN_KEY).(string)
	expireAt := payload.ExpiresAt.Time
	err := u.redisdb.Set(ctx, accessToken, middleware.BLACKLIST_CONTAINS_ACCESS_TOKEN, time.Until(expireAt)).Err()

	return err
}

// foget password
