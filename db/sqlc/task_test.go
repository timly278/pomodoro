package db

import (
	"context"
	"pomodoro/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createNewPomomTask(t *testing.T, user User) Task {

	params := CreateNewTaskParams{
		UserID:        user.ID,
		Content:       util.RandomString(15),
		Status:        int32(util.RandomInt(0, 2)),
		EstimatePomos: int32(util.RandomInt(1, 10)),
		ProgressPomos: 0,
	}
	newTask, err := testQueries.CreateNewTask(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newTask)

	return newTask
}

func TestCreateTask(t *testing.T) {
	user := createNewUser(t)
	createNewPomomTask(t, user)
}
