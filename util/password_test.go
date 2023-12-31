package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	pass := RandomString(6)
	hashedPassword1, err := HashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)
	err = VerifyPassword(pass, hashedPassword1)
	require.NoError(t, err)

	wrongPass := RandomString(6)
	err = VerifyPassword(wrongPass, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
