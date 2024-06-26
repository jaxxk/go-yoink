-- name: FollowFeed :one
INSERT INTO feed_follows (id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: Unfollow :exec
DELETE FROM feed_follows
WHERE feed_follows.id = $1;

-- name: GetAllFollowFeedsForUser :many
SELECT feed_follows.*
FROM feed_follows
INNER JOIN users ON users.id = feed_follows.user_id
WHERE users.id = $1;

-- name: GetFollowFeedByID :one
SELECT *
FROM feed_follows
WHERE feed_follows.id = $1;