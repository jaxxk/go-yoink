<<<<<<< HEAD
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
=======
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
>>>>>>> e5a6370aa5f3252380ec1f3e5f0dd89eb89f3ce4
