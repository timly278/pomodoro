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
SELECT t.goalperday, p.focus_degree, t.name as type_name, t.color as type_color, t.duration
FROM pomodoros p, types t 
WHERE t.id = p.type_id
AND (p.created_at::DATE) = sqlc.arg(query_date)::DATE AND p.user_id = $1
ORDER BY p.type_id
LIMIT $2
OFFSET $3;