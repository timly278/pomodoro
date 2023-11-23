// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CreateNewTask(ctx context.Context, arg CreateNewTaskParams) (Task, error)
	CreateNewType(ctx context.Context, arg CreateNewTypeParams) (Type, error)
	CreatePomodoroWithNoTask(ctx context.Context, arg CreatePomodoroWithNoTaskParams) (Pomodoro, error)
	CreatePomodoroWithTask(ctx context.Context, arg CreatePomodoroWithTaskParams) (Pomodoro, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetDaysAccessedInMonth(ctx context.Context, arg GetDaysAccessedInMonthParams) (int64, error)
	// todo: GetDaysAccessedInYear :one
	GetMinutesFocusedInMonth(ctx context.Context, arg GetMinutesFocusedInMonthParams) (int64, error)
	GetPomodoroByDate(ctx context.Context, arg GetPomodoroByDateParams) ([]GetPomodoroByDateRow, error)
	GetPomodoroByUserId(ctx context.Context, userID int64) ([]Pomodoro, error)
	GetTaskById(ctx context.Context, id int64) (Type, error)
	GetTotalDaysAccessed(ctx context.Context, userID int64) (int64, error)
	GetTypeById(ctx context.Context, id int64) (Type, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	ListTasks(ctx context.Context, userID int64) ([]Task, error)
	ListTypes(ctx context.Context, userID int64) ([]Type, error)
	UpdateTaskConfig(ctx context.Context, arg UpdateTaskConfigParams) (Task, error)
	UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) (Task, error)
	UpdateTypeById(ctx context.Context, arg UpdateTypeByIdParams) (Type, error)
	UpdateUserSetting(ctx context.Context, arg UpdateUserSettingParams) (User, error)
}

var _ Querier = (*Queries)(nil)
