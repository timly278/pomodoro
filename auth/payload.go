package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	SUBJECT_CLAIM_ACCESS_TOKEN  = "access token claim"
	SUBJECT_CLAIM_REFRESH_TOKEN = "refresh token claim"
)

type Payload struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func NewPayload(userid string, subject string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	issueAt := time.Now()
	expiredAt := time.Now().Add(duration)
	payload := Payload{
		tokenID,
		jwt.RegisteredClaims{
			ID:        userid,
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(issueAt),
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	return &payload, nil
}
