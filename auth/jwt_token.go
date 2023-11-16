package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenMaker struct {
	secretKey string
}

func NewJwtTokenMaker(secretKey string) *JwtTokenMaker {
	return &JwtTokenMaker{secretKey: secretKey}
}

// method to create token
func (maker *JwtTokenMaker) CreateToken(name string, duration time.Duration) (string, error) {
	payload, err := NewPayload(name, duration)
	if err != nil {
		return "", errors.New("cannot create token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}

// method to verify token
func (maker *JwtTokenMaker) VerifyToken(tokenString string) (*Payload, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrTokenUnverifiable
		}

		return []byte(maker.secretKey), nil
	}

	var payload Payload
	token, err := jwt.ParseWithClaims(tokenString, &payload, keyfunc)
	if err != nil {
		return nil, err
	}
	//test payload
	payload2, ok := token.Claims.(*Payload)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	fmt.Printf("payload: %v,\npayload2: %v\n", payload, payload2)
	return payload2, nil
}
