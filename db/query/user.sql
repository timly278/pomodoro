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

-- name: UpdateUser :one
UPDATE users
SET username = coalesce(sqlc.narg('username'), username),
    refresh_token = coalesce(sqlc.narg('refresh_token'), refresh_token),
    email_verified = coalesce(sqlc.narg('email_verified'), email_verified),
    hashed_password = coalesce(sqlc.narg('hashed_password'), hashed_password),
    password_changed_at = coalesce(sqlc.narg('password_changed_at'), password_changed_at),
    alarm_sound = coalesce(sqlc.narg('alarm_sound'), alarm_sound),
    repeat_alarm = coalesce(sqlc.narg('repeat_alarm'), repeat_alarm)
WHERE id = coalesce(sqlc.narg('id'), 0) OR email = sqlc.narg('email')
RETURNING *;
