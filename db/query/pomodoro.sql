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
SELECT t.id as type_id, p.focus_degree
FROM pomodoros p, types t 
WHERE t.id = p.type_id
AND (p.created_at::DATE) = sqlc.arg(query_date)::DATE AND p.user_id = $1
ORDER BY p.type_id
LIMIT $2
OFFSET $3;