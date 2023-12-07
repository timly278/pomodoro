-- name: CreatePomodoro :one
INSERT INTO pomodoros (
  user_id,
  type_id,
  task_id,
  focus_degree
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPomodoros :many
SELECT type_id, focus_degree, created_at::DATE
FROM pomodoros p 
WHERE (p.created_at::DATE) >= sqlc.arg(from_date)::DATE 
AND (p.created_at::DATE) <= sqlc.arg(to_date)::DATE 
AND p.user_id = $1
ORDER BY p.created_at::DATE, p.type_id
LIMIT $2
OFFSET $3;

-- name: GetDaysAccessed :one
SELECT count(*) FROM
(
  SELECT EXTRACT(day FROM created_at) as day FROM pomodoros 
  WHERE user_id = $1
  AND (p.created_at::DATE) >= sqlc.arg(from_date)::DATE 
  AND (p.created_at::DATE) <= sqlc.arg(to_date)::DATE 
  GROUP BY day
) as x LIMIT 1;

-- name: GetMinutesFocused :one
SELECT SUM(duration) as minutefocused FROM
(
  SELECT t.duration, p.created_at, p.type_id, p.user_id FROM 
	types as t, pomodoros as p
	WHERE t.id = p.type_id 
	AND p.user_id = $1 
  AND (p.created_at::DATE) >= sqlc.arg(from_date)::DATE 
  AND (p.created_at::DATE) <= sqlc.arg(to_date)::DATE 
) as x LIMIT 1;
