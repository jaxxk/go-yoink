-- +goose Up
ALTER TABLE feeds_users RENAME TO feed_follows;


-- +goose Down
ALTER TABLE feed_follows RENAME TO feeds_users;
