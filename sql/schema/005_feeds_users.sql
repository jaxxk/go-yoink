-- +goose Up
-- Step 1: Create feeds_users table
CREATE TABLE feeds_users (
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id VARCHAR(128) NOT NULL,
    feed_id VARCHAR(128) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
-- Step 1: Drop the feeds_users table
DROP TABLE feeds_users;