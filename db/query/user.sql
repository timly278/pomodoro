-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email,
  alarm_sound,
  repeat_alarm
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

--TODO: change password