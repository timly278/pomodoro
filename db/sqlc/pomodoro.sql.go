// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: pomodoro.sql

package db

import (
	"context"
)

const createPomodoro = `-- name: CreatePomodoro :one
INSERT INTO pomodoros (
  user_id,
  type_id,
  task_id,
  focus_degree
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, type_id, task_id, focus_degree, created_at
`

type CreatePomodoroParams struct {
	UserID      int64 `json:"user_id"`
	TypeID      int32 `json:"type_id"`
	TaskID      int64 `json:"task_id"`
	FocusDegree int32 `json:"focus_degree"`
}

func (q *Queries) CreatePomodoro(ctx context.Context, arg CreatePomodoroParams) (Pomodoro, error) {
	row := q.db.QueryRowContext(ctx, createPomodoro,
		arg.UserID,
		arg.TypeID,
		arg.TaskID,
		arg.FocusDegree,
	)
	var i Pomodoro
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TypeID,
		&i.TaskID,
		&i.FocusDegree,
		&i.CreatedAt,
	)
	return i, err
}

const getPomodoroByUserId = `-- name: GetPomodoroByUserId :many
SELECT id, user_id, type_id, task_id, focus_degree, created_at FROM pomodoros
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetPomodoroByUserId(ctx context.Context, userID int64) ([]Pomodoro, error) {
	rows, err := q.db.QueryContext(ctx, getPomodoroByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pomodoro{}
	for rows.Next() {
		var i Pomodoro
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TypeID,
			&i.TaskID,
			&i.FocusDegree,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
