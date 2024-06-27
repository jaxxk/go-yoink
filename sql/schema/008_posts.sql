-- +goose Up
-- table for posts, a single entry from feeds
CREATE TABLE posts(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    title VARCHAR(32),
    url VARCHAR(128) NOT NULL,
    description VARCHAR(128),
    published_at TIMESTAMP NOT NULL, 
    feed_id VARCHAR(128),
    CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
-- drops posts table
DROP TABLE posts;