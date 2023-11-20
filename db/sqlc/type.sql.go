// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: type.sql

package db

import (
	"context"
)

const createNewType = `-- name: CreateNewType :one
INSERT INTO types (
  user_id,
  name,
  color,
  duration,
  shortbreak,
  longbreak,
  longbreakinterval,
  autostart_break
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING id, user_id, name, color, duration, shortbreak, longbreak, longbreakinterval, autostart_break
`

type CreateNewTypeParams struct {
	UserID            int64  `json:"user_id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	Duration          int32  `json:"duration"`
	Shortbreak        int32  `json:"shortbreak"`
	Longbreak         int32  `json:"longbreak"`
	Longbreakinterval int32  `json:"longbreakinterval"`
	AutostartBreak    bool   `json:"autostart_break"`
}

func (q *Queries) CreateNewType(ctx context.Context, arg CreateNewTypeParams) (Type, error) {
	row := q.db.QueryRowContext(ctx, createNewType,
		arg.UserID,
		arg.Name,
		arg.Color,
		arg.Duration,
		arg.Shortbreak,
		arg.Longbreak,
		arg.Longbreakinterval,
		arg.AutostartBreak,
	)
	var i Type
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Color,
		&i.Duration,
		&i.Shortbreak,
		&i.Longbreak,
		&i.Longbreakinterval,
		&i.AutostartBreak,
	)
	return i, err
}

const getTypeById = `-- name: GetTypeById :one
SELECT id, user_id, name, color, duration, shortbreak, longbreak, longbreakinterval, autostart_break FROM types
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTypeById(ctx context.Context, id int32) (Type, error) {
	row := q.db.QueryRowContext(ctx, getTypeById, id)
	var i Type
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Color,
		&i.Duration,
		&i.Shortbreak,
		&i.Longbreak,
		&i.Longbreakinterval,
		&i.AutostartBreak,
	)
	return i, err
}

const listTypes = `-- name: ListTypes :many
SELECT id, user_id, name, color, duration, shortbreak, longbreak, longbreakinterval, autostart_break FROM types
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) ListTypes(ctx context.Context, userID int64) ([]Type, error) {
	rows, err := q.db.QueryContext(ctx, listTypes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Type{}
	for rows.Next() {
		var i Type
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Color,
			&i.Duration,
			&i.Shortbreak,
			&i.Longbreak,
			&i.Longbreakinterval,
			&i.AutostartBreak,
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

const updateTypeById = `-- name: UpdateTypeById :one
UPDATE types
SET name = $2,
    color = $3,
    shortbreak = $4,
    duration = $5,
    longbreak = $6,
    longbreakinterval = $7,
    autostart_break = $8
WHERE id = $1
RETURNING id, user_id, name, color, duration, shortbreak, longbreak, longbreakinterval, autostart_break
`

type UpdateTypeByIdParams struct {
	ID                int32  `json:"id"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	Shortbreak        int32  `json:"shortbreak"`
	Duration          int32  `json:"duration"`
	Longbreak         int32  `json:"longbreak"`
	Longbreakinterval int32  `json:"longbreakinterval"`
	AutostartBreak    bool   `json:"autostart_break"`
}

func (q *Queries) UpdateTypeById(ctx context.Context, arg UpdateTypeByIdParams) (Type, error) {
	row := q.db.QueryRowContext(ctx, updateTypeById,
		arg.ID,
		arg.Name,
		arg.Color,
		arg.Shortbreak,
		arg.Duration,
		arg.Longbreak,
		arg.Longbreakinterval,
		arg.AutostartBreak,
	)
	var i Type
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Color,
		&i.Duration,
		&i.Shortbreak,
		&i.Longbreak,
		&i.Longbreakinterval,
		&i.AutostartBreak,
	)
	return i, err
}