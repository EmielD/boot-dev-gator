-- +goose Up
CREATE TABLE posts(
    id UUID not null PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    title varchar(255) not null,
    url varchar(255) UNIQUE not null,
    description varchar(255),
    published_at TIMESTAMP,
    feed_id UUID not null
);

-- +goose Down
DROP TABLE posts;