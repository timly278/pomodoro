-- name: GetTotalDaysAccessed :one
SELECT count(*) as count FROM
(
  SELECT (created_at::DATE) as days
  FROM pomodoros where user_id = $1 
  GROUP BY (created_at::DATE)
) as x LIMIT 1;

-- name: GetDaysAccessedInMonth :one
SELECT count(*) FROM
(
    SELECT EXTRACT(day FROM created_at) as day FROM pomodoros 
    WHERE user_id = $1
    AND EXTRACT(MONTH FROM created_at) = sqlc.arg(month_id)::int
    AND EXTRACT(YEAR FROM created_at) = sqlc.arg(year)::int
    GROUP BY day
) as x LIMIT 1;
-- todo: GetDaysAccessedInYear :one

-- name: GetMinutesFocusedInMonth :one
SELECT SUM(duration) as minutefocused FROM
(
    SELECT t.duration, p.created_at, p.type_id, p.user_id FROM 
	types as t, pomodoros as p
	WHERE t.id = p.type_id 
	AND p.user_id = $1 
	AND EXTRACT(MONTH FROM created_at) = sqlc.arg(month_id)::int
    AND EXTRACT(YEAR FROM created_at) = sqlc.arg(year)::int
) as x LIMIT 1;
