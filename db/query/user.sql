-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUserSetting :one
UPDATE users
SET   username = $2,
      alarm_sound = $3,
      repeat_alarm = $4
WHERE id = $1
RETURNING *;

--TODO: change password