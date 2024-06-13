-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, api_key
FROM users
WHERE api_key = $1;