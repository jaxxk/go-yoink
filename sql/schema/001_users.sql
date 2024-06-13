-- +goose Up
CREATE TABLE users(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(30) NOT NULL
);

-- +goose Down
DROP TABLE users;
