-- +goose Up
-- table for posts, a single entry from feeds
CREATE TABLE posts(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    title VARCHAR(128) NOT NULL,
    url VARCHAR(128) NOT NULL,
    description VARCHAR(800) NOT NULL,
    published_at TIMESTAMP NOT NULL, 
    feed_id VARCHAR(128) NOT NULL,
    CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
-- drops posts table
DROP TABLE posts;