-- name: FollowFeed :one
INSERT INTO feeds_users (id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: Unfollow :exec
DELETE FROM feeds_users
WHERE feeds_users.id = $1;