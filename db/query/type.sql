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
SET name = $3,
    color = $4,
    goalperday = $5,
    shortbreak = $6,
    duration = $7,
    longbreak = $8,
    longbreakinterval = $9,
    autostart_break = $10
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: GetTypeById :one
SELECT * FROM types
WHERE id = $1 AND user_id = $2 LIMIT 1;