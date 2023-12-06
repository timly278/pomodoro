package logging

import (
	"context"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/security"
	"pomodoro/shared/response"
	"pomodoro/util"
	"strconv"
)

const (
	REFRESH_TOKEN_FAKE = ""
)

type TokenLogging struct {
	store      db.Store
	tokenMaker security.TokenMaker
	conf       *util.Config
}

func NewTokenLogging(store db.Store, tokenMaker security.TokenMaker, conf *util.Config) *TokenLogging {
	return &TokenLogging{
		store:      store,
		tokenMaker: tokenMaker,
		conf:       conf,
	}
}

func (t *TokenLogging) RefreshTokens(ctx context.Context, req delivery.RefreshTokenRequest) (*response.NewTokensResponse, error) {
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

func (t *TokenLogging) newTokens(ctx context.Context, userId int64) (rsp *response.NewTokensResponse, err error) {
	accessToken, err := newAccessToken(t.tokenMaker, t.conf, userId)
	if err != nil {
		return
	}

	refreshToken, err := newRefreshToken(t.tokenMaker, t.conf, userId)
	if err != nil {
		return
	}

	params := db.UpdateRefreshTokenParams{
		ID:           userId,
		RefreshToken: refreshToken,
	}
	_, err = t.store.UpdateRefreshToken(ctx, params)
	if err != nil {
		return
	}

	rsp = &response.NewTokensResponse{
		RefreshToken: refreshToken,
		RTExpireIn:   int64(t.conf.RefreshTokenDuration.Seconds()),
		AccessToken:  accessToken,
		ATExpireIn:   int64(t.conf.AccessTokenDuration.Seconds()),
	}
	return
}

func (t *TokenLogging) validateRefreshToken(ctx context.Context, rToken string) (userid int64, err error) {

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
		return
	}

	return int64(userId), nil
}

func (t *TokenLogging) revokeRefreshToken(ctx context.Context, userId int64) (err error) {
	_, err = t.store.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
		ID:           userId,
		RefreshToken: REFRESH_TOKEN_FAKE,
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
