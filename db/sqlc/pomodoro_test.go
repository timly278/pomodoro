package db

import (
	"context"
	"database/sql"
	"pomodoro/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createPomoWithTask(t *testing.T, userID, typeID, taskID64 int64) Pomodoro {
	var taskID sql.NullInt64
	err := taskID.Scan(taskID64)
	require.NoError(t, err)
	params := CreatePomodoroWithTaskParams{
		UserID:      userID,
		TypeID:      typeID,
		TaskID:      taskID,
		FocusDegree: int32(util.RandomInt(1, 5)),
	}
	newpomo, err := testQueries.CreatePomodoroWithTask(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newpomo)
	require.Equal(t, userID, newpomo.UserID)
	require.Equal(t, typeID, newpomo.TypeID)
	require.Equal(t, newpomo.FocusDegree, newpomo.FocusDegree)
	require.False(t, newpomo.CreatedAt.IsZero())

	taskIdValue, _ := newpomo.TaskID.Value()
	require.Equal(t, taskID64, taskIdValue)
	return newpomo
}

func TestCreatePomodoroWithTask(t *testing.T) {
	user := createNewUser(t)
	pomotype := createNewPomoType(t, user)
	pomotask := createNewPomomTask(t, user)
	createPomoWithTask(t, user.ID, pomotype.ID, pomotask.ID)
}

func createPomoWithNoTask(t *testing.T, userID, typeID int64) Pomodoro {
	params := CreatePomodoroWithTaskParams{
		UserID:      userID,
		TypeID:      typeID,
		FocusDegree: int32(util.RandomInt(1, 5)),
	}
	newpomo, err := testQueries.CreatePomodoroWithTask(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newpomo)
	require.Equal(t, userID, newpomo.UserID)
	require.Equal(t, typeID, newpomo.TypeID)
	require.Equal(t, newpomo.FocusDegree, newpomo.FocusDegree)
	require.False(t, newpomo.CreatedAt.IsZero())

	taskIdValue, _ := newpomo.TaskID.Value()
	require.Equal(t, nil, taskIdValue)

	return newpomo
}
func TestCreatePomodoroWitNoTask(t *testing.T) {
	user := createNewUser(t)
	pomotype := createNewPomoType(t, user)
	createPomoWithNoTask(t, user.ID, pomotype.ID)
}
func TestGetPomodoroByDate(t *testing.T) {
	user := createNewUser(t)
	pomotype := createNewPomoType(t, user)
	n := 5

	var pomodoros []Pomodoro
	for i := 0; i < n; i++ {
		pomodoros = append(pomodoros, createPomoWithNoTask(t, user.ID, pomotype.ID))
	}

	params := GetPomodoroByDateParams{
		UserID:    user.ID,
		Limit:     int32(n),
		Offset:    0,
		QueryDate: time.Now(),
	}
	pomoRows, err := testQueries.GetPomodoroByDate(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, n, len(pomoRows))
	for i := 0; i < n; i++ {
		require.Equal(t, pomotype.ID, pomoRows[i].TypeID)
		require.Equal(t, pomodoros[i].FocusDegree, pomoRows[i].FocusDegree)
	}
}
