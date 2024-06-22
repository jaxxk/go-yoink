-- name: CreateFeed :one
INSERT INTO feeds (id,created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * 
FROM feeds;

-- name: GetFeedByUserID :one
SELECT * 
FROM feeds
WHERE feeds.user_id = $1;