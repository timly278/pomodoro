-- name: CreateNewTask :one
INSERT INTO tasks (
  user_id,
  content,
  status,
  estimate_pomos,
  progress_pomos
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks
WHERE user_id = $1
ORDER BY id;

-- name: UpdateTaskConfig :one
UPDATE tasks
SET     content = $3,
        estimate_pomos = $4
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET     status = $3,
        progress_pomos = $4
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: GetTaskById :one
SELECT * FROM types
WHERE id = $1 AND user_id = $2 LIMIT 1;
