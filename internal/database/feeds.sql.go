// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feeds.sql

package database

import (
	"context"
	"time"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id,created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING name, url, user_id, id, created_at, updated_at
`

type CreateFeedParams struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllFeeds = `-- name: GetAllFeeds :many
SELECT name, url, user_id, id, created_at, updated_at 
FROM feeds
`

func (q *Queries) GetAllFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getAllFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getFeedByUserID = `-- name: GetFeedByUserID :one
SELECT name, url, user_id, id, created_at, updated_at 
FROM feeds
WHERE feeds.user_id = $1
`

func (q *Queries) GetFeedByUserID(ctx context.Context, userID string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByUserID, userID)
	var i Feed
	err := row.Scan(
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
