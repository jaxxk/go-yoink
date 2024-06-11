<<<<<<< HEAD
-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE users;
=======
-- +goose Up
CREATE TABLE users(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(30) NOT NULL
);

-- +goose Down
DROP TABLE users;
>>>>>>> e5a6370aa5f3252380ec1f3e5f0dd89eb89f3ce4
