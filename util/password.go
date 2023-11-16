package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates hashed password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashed), nil
}

// VerifyPassword verifies password
func VerifyPassword(password, hashedpassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err
}
