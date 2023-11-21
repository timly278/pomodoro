-- name: CreatePomodoroWithTask :one
INSERT INTO pomodoros (
  user_id,
  type_id,
  task_id,
  focus_degree
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: CreatePomodoroWithNoTask :one
INSERT INTO pomodoros (
  user_id,
  type_id,
  focus_degree
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPomodoroByUserId :many
SELECT * FROM pomodoros
WHERE user_id = $1 LIMIT 1;

-- name: GetPomodoroByDate :many
SELECT * FROM pomodoros
WHERE (created_at::DATE) = sqlc.arg(CreatedDate)::DATE AND user_id = $1;

