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
SET     content = $2,
        estimate_pomos = $3
WHERE id = $1
RETURNING *;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET     status = $2,
        progress_pomos = $3
WHERE id = $1
RETURNING *;

-- name: GetTaskById :one
SELECT * FROM types
WHERE id = $1 LIMIT 1;
