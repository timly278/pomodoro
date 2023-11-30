-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

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

-- name: UpdateRefreshToken :one
UPDATE users
SET refresh_token = $2
WHERE id = $1
RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE users
SET email_verified = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET hashed_password = $2,
    password_changed_at = $3
WHERE id = $1
RETURNING *;