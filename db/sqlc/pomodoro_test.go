package db

import (
	"context"
	"database/sql"
	"fmt"
	"pomodoro/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createPomo(t *testing.T, userID, typeID, taskID64 int64) Pomodoro {
	params := CreatePomodoroParams{
		UserID: userID,
		TypeID: typeID,
		TaskID: sql.NullInt64{
			Int64: taskID64,
			Valid: bool(taskID64 != 0),
		},
		FocusDegree: int32(util.RandomInt(1, 5)),
	}
	fmt.Println(userID, typeID, params.TaskID)
	newpomo, err := testQueries.CreatePomodoro(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, newpomo)
	require.Equal(t, userID, newpomo.UserID)
	require.Equal(t, typeID, newpomo.TypeID)
	require.Equal(t, newpomo.FocusDegree, newpomo.FocusDegree)
	require.False(t, newpomo.CreatedAt.IsZero())

	fmt.Printf("newpomo.TaskID = %v\n", newpomo.TaskID)
	return newpomo
}

func TestCreatePomodoroWithTask(t *testing.T) {
	user := createNewUser(t)
	user.ID = 49
	pomotype := createNewPomoType(t, user)
	pomotask := createNewPomomTask(t, user)
	newpomo := createPomo(t, user.ID, pomotype.ID, pomotask.ID)
	require.Equal(t, pomotask.ID, newpomo.TaskID.Int64)
}

func TestCreatePomodoroWitNoTask(t *testing.T) {
	user := createNewUser(t)
	user.ID = 49
	pomotype := createNewPomoType(t, user)
	newpomo := createPomo(t, user.ID, pomotype.ID, 0)
	fmt.Printf("newpomo = %v\n", newpomo)
	require.Equal(t, int64(0), newpomo.TaskID.Int64)
}
