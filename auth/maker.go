package auth

import "time"


type TokenMaker interface {
	// method to create token
	CreateToken(name string, duration time.Duration) (string, error)

	// method to verify token
	VerifyToken(token string) (*Payload, error)
}