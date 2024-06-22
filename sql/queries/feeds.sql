-- name: CreateFeed :one
INSERT INTO feeds (id,created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * 
FROM feeds;

-- name: FollowFeed :one
INSERT INTO feeds_users (id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;