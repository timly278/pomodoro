-- name: CreateNewGoal :one
INSERT INTO goalperday (
  user_id,
  pomonum,
  type_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListGoals :many
SELECT * FROM goalperday
WHERE user_id = $1
ORDER BY id;

-- name: UpdateGoal :one
UPDATE goalperday
SET    
    pomonum = $2,
    type_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteGoal :exec
DELETE FROM goalperday
WHERE id = $1;