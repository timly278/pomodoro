package db

import (
	"context"
	"pomodoro/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createNewPomoType(t *testing.T, user User) Type {
	params := CreateNewTypeParams{
		UserID:            user.ID,
		Name:              util.RandomString(6),
		Color:             util.RandomColor(),
		Duration:          int32(util.RandomInt(10, 45)),
		Shortbreak:        int32(util.RandomInt(2, 10)),
		Longbreak:         int32(util.RandomInt(5, 20)),
		Longbreakinterval: int32(util.RandomInt(1, 5)),
		AutostartBreak:    true,
	}
	newType, err := testQueries.CreateNewType(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newType)
	require.Equal(t, user.ID, newType.UserID)
	require.Equal(t, params.Name, newType.Name)
	require.Equal(t, params.Color, newType.Color)
	require.Equal(t, params.Duration, newType.Duration)
	require.Equal(t, params.Shortbreak, newType.Shortbreak)
	require.Equal(t, params.Longbreak, newType.Longbreak)
	require.Equal(t, params.Longbreakinterval, newType.Longbreakinterval)
	require.Equal(t, params.AutostartBreak, newType.AutostartBreak)

	return newType
}

func TestCreateNewType(t *testing.T) {
	user := createNewUser(t)
	createNewPomoType(t, user)
}

func TestListTypes(t *testing.T) {
	user := createNewUser(t)
	numOfType := 5
	newTypes := make([]Type, numOfType)

	for i := 0; i < numOfType; i++ {
		newTypes[i] = createNewPomoType(t, user)
	}

	types, err := testQueries.ListTypes(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, types)
	for i := 0; i < numOfType; i++ {
		require.NotEmpty(t, types[i])
		require.NotZero(t, types[i].ID)
		require.Equal(t, user.ID, types[i].UserID)
	}
}
