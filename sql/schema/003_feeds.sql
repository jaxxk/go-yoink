-- +goose Up
CREATE TABLE feeds (
    name VARCHAR(32) NOT NULL UNIQUE PRIMARY KEY,
    url VARCHAR(2000) NOT NULL,
    user_id VARCHAR(128) NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
);

-- +goose Down
DROP TABLE feeds;