package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const MIN_SECRETKEY_SIZE = 32

type JwtTokenMaker struct {
	secretKey string
}

func NewJwtTokenMaker(secretKey string) (*JwtTokenMaker, error) {
	if len(secretKey) < MIN_SECRETKEY_SIZE {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", MIN_SECRETKEY_SIZE)
	}

	return &JwtTokenMaker{secretKey: secretKey}, nil
}

// method to create token
func (maker *JwtTokenMaker) CreateToken(id, subject string, duration time.Duration) (string, error) {
	payload, err := NewPayload(id, subject, duration)
	if err != nil {
		return "", errors.New("cannot create token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}

// method to verify token
func (maker *JwtTokenMaker) VerifyToken(tokenString, subject string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrTokenUnverifiable
		}

		return []byte(maker.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyfunc)
	if err != nil {
		return nil, err
	}
	payload2, ok := token.Claims.(*Payload)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	if subject != payload2.Subject {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return payload2, nil
}
