// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: follow_feeds.sql

package database

import (
	"context"
	"time"
)

const followFeed = `-- name: FollowFeed :one
INSERT INTO feed_follows (id,created_at,updated_at,user_id,feed_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING id, created_at, updated_at, user_id, feed_id
`

type FollowFeedParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	FeedID    string
}

func (q *Queries) FollowFeed(ctx context.Context, arg FollowFeedParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, followFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const getAllFollowFeedsForUser = `-- name: GetAllFollowFeedsForUser :many
SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.user_id, feed_follows.feed_id
FROM feed_follows
INNER JOIN users ON users.id = feed_follows.user_id
WHERE users.id = $1
`

func (q *Queries) GetAllFollowFeedsForUser(ctx context.Context, id string) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getAllFollowFeedsForUser, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowFeedByID = `-- name: GetFollowFeedByID :one
SELECT id, created_at, updated_at, user_id, feed_id
FROM feed_follows
WHERE feed_follows.id = $1
`

func (q *Queries) GetFollowFeedByID(ctx context.Context, id string) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, getFollowFeedByID, id)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const unfollow = `-- name: Unfollow :exec
DELETE FROM feed_follows
WHERE feed_follows.id = $1
`

func (q *Queries) Unfollow(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, unfollow, id)
	return err
}