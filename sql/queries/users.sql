-- name: CreateUser :one
INSERT INTO users(name, updated_at)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE users.name = $1;

-- name: DeleteAllUsers :exec
TRUNCATE TABLE users;

-- name: GetAllUsers :many
SELECT * FROM users;
