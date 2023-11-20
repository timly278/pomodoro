package db

import (
	"context"
	"pomodoro/util"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TEST_DEFAULT_REPEATALARM = 2
)

func createNewUser(t *testing.T) User {
	userParams := CreateUserParams{
		Username:       util.RandomString(6),
		HashedPassword: util.RandomString(12),
		Email:          util.RandomString(3) + "@gmail.com",
		AlarmSound:     util.RandomAlarmSound(),
		RepeatAlarm:    TEST_DEFAULT_REPEATALARM,
	}
	user, err := testQueries.CreateUser(context.Background(), userParams)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, userParams.Username, user.Username)
	require.Equal(t, userParams.HashedPassword, user.HashedPassword)
	require.Equal(t, userParams.Email, user.Email)
	require.Equal(t, userParams.AlarmSound, user.AlarmSound)
	require.Equal(t, userParams.RepeatAlarm, user.RepeatAlarm)
	require.Equal(t, userParams.RepeatAlarm, user.RepeatAlarm)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createNewUser(t)
}
