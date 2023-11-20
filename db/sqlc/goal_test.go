package db

import (
	"context"
	"pomodoro/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createNewPomoGoal(t *testing.T) Goalperday {
	user := createNewUser(t)
	pomoType := createNewPomoType(t, user)
	params := CreateNewGoalParams{
		UserID:  user.ID,
		Pomonum: int32(util.RandomInt(0, 10)),
		TypeID:  pomoType.ID,
	}
	newGoal, err := testQueries.CreateNewGoal(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newGoal)
	require.Equal(t, user.ID, newGoal.UserID)
	require.Equal(t, pomoType.ID, newGoal.TypeID)
	require.Equal(t, params.Pomonum, newGoal.Pomonum)

	return newGoal
}

func TestCreateNewGoal(t *testing.T) {
	newGoal := createNewPomoGoal(t)
	params := CreateNewGoalParams{
		UserID:  newGoal.ID,
		Pomonum: int32(util.RandomInt(0, 10)),
		TypeID:  newGoal.TypeID,
	}
	newGoalerr, err := testQueries.CreateNewGoal(context.Background(), params)
	require.Error(t, err)
	require.Empty(t, newGoalerr)
}
