-- name: ReadAll :many
SELECT * FROM todos
ORDER BY completed ASC, created_at ASC;

-- name: ReadOne :one
SELECT * FROM todos
WHERE id = ?;

-- name: Create :one
INSERT INTO todos (
    title
) VALUES (?)
RETURNING *;

-- name: ChangeStatus :one
UPDATE todos
SET completed = ?
WHERE id = ?
RETURNING *;

-- name: Delete :exec
DELETE FROM todos
WHERE id = ?;
