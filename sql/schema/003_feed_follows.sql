-- +goose Up
CREATE TABLE feed_follows(
    id UUID not null PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    feed_id UUID not null references feeds ON DELETE CASCADE,
    user_id UUID not null references users ON DELETE CASCADE,
    UNIQUE(feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;