-- name: CreateNewType :one
INSERT INTO types (
  user_id,
  name,
  color,
  goalperday,
  duration,
  shortbreak,
  longbreak,
  longbreakinterval,
  autostart_break
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: ListTypes :many
SELECT * FROM types
WHERE user_id = $1
ORDER BY id;

-- name: UpdateTypeById :one
UPDATE types
SET name = $2,
    color = $3,
    goalperday = $4,
    shortbreak = $5,
    duration = $6,
    longbreak = $7,
    longbreakinterval = $8,
    autostart_break = $9
WHERE id = $1
RETURNING *;

-- name: GetTypeById :one
SELECT * FROM types
WHERE id = $1 LIMIT 1;