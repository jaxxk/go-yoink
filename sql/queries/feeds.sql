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

-- name: GetNextFeedsToFetch :many
SELECT name, url
FROM feeds
ORDER BY last_fetched_at IS NULL DESC, last_fetched_at ASC
LIMIT $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $1, last_fetched_at = $1
WHERE user_id = $2;