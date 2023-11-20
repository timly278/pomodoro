-- name: CreateNewType :one
INSERT INTO types (
  user_id,
  name,
  color,
  duration,
  shortbreak,
  longbreak,
  longbreakinterval,
  autostart_break
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: ListTypes :many
SELECT * FROM types
WHERE user_id = $1
ORDER BY id;

-- name: UpdateTypeById :one
UPDATE types
SET name = $2,
    color = $3,
    shortbreak = $4,
    duration = $5,
    longbreak = $6,
    longbreakinterval = $7,
    autostart_break = $8
WHERE id = $1
RETURNING *;

-- name: GetTypeById :one
SELECT * FROM types
WHERE id = $1 LIMIT 1;