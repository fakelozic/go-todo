-- name: CreateTask :one
INSERT INTO tasks (id, task, status, created_at, updated_at, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllTasks :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1 AND user_id = $2;

-- name: UpdateTask :one
UPDATE tasks
SET
    task = $1,
    updated_at = $2
    WHERE id = $3 AND user_id = $4
RETURNING *;

-- name: ToggleTask :one
UPDATE tasks
SET
    status = $1,
    updated_at = $2
    WHERE id = $3 AND user_id = $4
RETURNING *;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = $1 AND user_id = $2
RETURNING *;