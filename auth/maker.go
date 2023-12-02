package auth

import "time"

type TokenMaker interface {
	// method to create token
	CreateToken(id, subject string, duration time.Duration) (string, error)

	// method to verify token
	VerifyToken(token, subject string) (*Payload, error)
}
