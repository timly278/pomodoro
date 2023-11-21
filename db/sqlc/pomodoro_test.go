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
	for i := 0; i < n; i++ {
		createPomoWithNoTask(t, user.ID, pomotype.ID)
	}

	params := GetPomodoroByDateParams{
		UserID:      user.ID,
		Createddate: time.Now(),
	}
	pomodoros, err := testQueries.GetPomodoroByDate(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, n, len(pomodoros))
	for i := 0; i < n; i++ {
		yearp, monthp, dayp := params.Createddate.Date()
		year, month, day := pomodoros[i].CreatedAt.Date()

		require.Equal(t, yearp, year)
		require.Equal(t, monthp, month)
		require.Equal(t, dayp, day)
	}
}

// FormCreatedDate was deprecated
// func FormCreatedDate(year int, month time.Month, day int) string {
// 	dateString := make([]byte, 1)
// 	dateString = append(dateString, []byte(strconv.Itoa(year))...)
// 	dateString = append(dateString, '-')
// 	dateString = append(dateString, []byte(strconv.Itoa(int(month)))...)
// 	dateString = append(dateString, '-')
// 	dateString = append(dateString, []byte(strconv.Itoa(day))...)
// 	return string(dateString)
// }
