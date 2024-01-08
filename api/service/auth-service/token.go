package authservice

import (
	"context"
	"database/sql"
	"errors"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/shared/response"
	"pomodoro/util"
	"strconv"
)

const (
	REVOKED_REFRESH_TOKEN = ""
)

func (t *authService) RefreshTokens(ctx context.Context, req delivery.RefreshTokenRequest) (*response.NewTokensResponse, error) {
	userId, err := t.validateRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	newTokens, err := t.newTokens(ctx, userId)
	if err != nil {
		return nil, err
	}

	return newTokens, nil
}

func (t *authService) newTokens(ctx context.Context, userId int64) (*response.NewTokensResponse, error) {
	accessToken, err := newAccessToken(t.tokenMaker, t.config, userId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := newRefreshToken(t.tokenMaker, t.config, userId)
	if err != nil {
		return nil, err
	}

	params := db.UpdateUserParams{
		ID: sql.NullInt64{
			Int64: userId,
			Valid: true,
		},
		RefreshToken: sql.NullString{
			String: refreshToken,
			Valid:  true,
		},
		IsBlocked: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
	}
	_, err = t.store.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	rsp := &response.NewTokensResponse{
		RefreshToken: refreshToken,
		RTExpireIn:   int64(t.config.RefreshTokenDuration.Seconds()),
		AccessToken:  accessToken,
		ATExpireIn:   int64(t.config.AccessTokenDuration.Seconds()),
	}
	return rsp, nil
}

func (t *authService) validateRefreshToken(ctx context.Context, rToken string) (userid int64, err error) {

	// verify RefreshToken
	payload, err := t.tokenMaker.VerifyToken(rToken, security.SUBJECT_CLAIM_REFRESH_TOKEN)
	if err != nil {
		return 0, err
	}

	userId, err := payload.GetIntegerUserID()
	if err != nil {
		return 0, err
	}

	user, err := t.store.GetUserById(ctx, int64(userId))
	if err != nil {
		return 0, err
	}

	if user.RefreshToken != rToken {
		err = t.revokeRefreshToken(ctx, user.ID)
		if err != nil {
			return 0, err
		}
		return 0, errors.New("detected reusing fresh-token")
	}

	return int64(userId), nil
}

// revokeRefreshToken revokes refresh_token and log out the user from the app
func (t *authService) revokeRefreshToken(ctx context.Context, userId int64) (err error) {
	_, err = t.store.UpdateUser(ctx, db.UpdateUserParams{
		ID: sql.NullInt64{
			Int64: userId,
			Valid: true,
		},
		RefreshToken: sql.NullString{
			String: REVOKED_REFRESH_TOKEN,
			Valid:  true,
		},
		IsBlocked: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	})
	return
}

func newRefreshToken(tokenMaker security.TokenMaker, conf *util.Config, userId int64) (refreshToken string, err error) {
	return tokenMaker.CreateToken(
		strconv.FormatInt(userId, 10),
		security.SUBJECT_CLAIM_REFRESH_TOKEN,
		conf.RefreshTokenDuration)
}
func newAccessToken(tokenMaker security.TokenMaker, conf *util.Config, userId int64) (refreshToken string, err error) {
	return tokenMaker.CreateToken(
		strconv.FormatInt(userId, 10),
		security.SUBJECT_CLAIM_ACCESS_TOKEN,
		conf.AccessTokenDuration)
}

// func storeRefreshToken(ctx context.Context, store db.Store, params *db.UpdateRefreshTokenParams) error {
// 	_, err := store.UpdateRefreshToken(ctx, *params)
// 	return err
// }
