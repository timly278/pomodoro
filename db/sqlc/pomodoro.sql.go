// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: pomodoro.sql

package db

import (
	"context"
	"database/sql"
	"time"
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
	UserID      int64         `json:"user_id"`
	TypeID      int64         `json:"type_id"`
	TaskID      sql.NullInt64 `json:"task_id"`
	FocusDegree int32         `json:"focus_degree"`
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

const getDaysAccessed = `-- name: GetDaysAccessed :one
SELECT count(*) FROM
(
  SELECT EXTRACT(day FROM created_at) as day FROM pomodoros 
  WHERE user_id = $1
  AND (p.created_at::DATE) >= $2::DATE 
  AND (p.created_at::DATE) <= $3::DATE 
  GROUP BY day
) as x LIMIT 1
`

type GetDaysAccessedParams struct {
	UserID   int64     `json:"user_id"`
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
}

func (q *Queries) GetDaysAccessed(ctx context.Context, arg GetDaysAccessedParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getDaysAccessed, arg.UserID, arg.FromDate, arg.ToDate)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMinutesFocused = `-- name: GetMinutesFocused :one
SELECT SUM(duration) as minutefocused FROM
(
  SELECT t.duration, p.created_at, p.type_id, p.user_id FROM 
	types as t, pomodoros as p
	WHERE t.id = p.type_id 
	AND p.user_id = $1 
  AND (p.created_at::DATE) >= $2::DATE 
  AND (p.created_at::DATE) <= $3::DATE 
) as x LIMIT 1
`

type GetMinutesFocusedParams struct {
	UserID   int64     `json:"user_id"`
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
}

func (q *Queries) GetMinutesFocused(ctx context.Context, arg GetMinutesFocusedParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getMinutesFocused, arg.UserID, arg.FromDate, arg.ToDate)
	var minutefocused int64
	err := row.Scan(&minutefocused)
	return minutefocused, err
}

const getPomodoros = `-- name: GetPomodoros :many
SELECT type_id, focus_degree, created_at::DATE
FROM pomodoros p 
WHERE (p.created_at::DATE) >= $4::DATE 
AND (p.created_at::DATE) <= $5::DATE 
AND p.user_id = $1
ORDER BY p.created_at::DATE, p.type_id
LIMIT $2
OFFSET $3
`

type GetPomodorosParams struct {
	UserID   int64     `json:"user_id"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
}

type GetPomodorosRow struct {
	TypeID      int64     `json:"type_id"`
	FocusDegree int32     `json:"focus_degree"`
	CreatedAt   time.Time `json:"created_at"`
}

func (q *Queries) GetPomodoros(ctx context.Context, arg GetPomodorosParams) ([]GetPomodorosRow, error) {
	rows, err := q.db.QueryContext(ctx, getPomodoros,
		arg.UserID,
		arg.Limit,
		arg.Offset,
		arg.FromDate,
		arg.ToDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPomodorosRow{}
	for rows.Next() {
		var i GetPomodorosRow
		if err := rows.Scan(&i.TypeID, &i.FocusDegree, &i.CreatedAt); err != nil {
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
