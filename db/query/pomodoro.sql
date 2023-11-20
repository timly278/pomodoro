-- name: CreatePomodoro :one
INSERT INTO pomodoros (
  user_id,
  type_id,
  task_id,
  focus_degree
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPomodoroByUserId :many
SELECT * FROM pomodoros
WHERE user_id = $1 LIMIT 1;

-- todo: GetPomodoroByDate, Week, month? along with type? and goal?
