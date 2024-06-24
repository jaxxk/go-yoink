-- +goose Up
-- track when the last fetch was
ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP NULL;


-- +goose Down
ALTER TABLE feeds
DROP COLUMN LAST_FETCHED_AT;
