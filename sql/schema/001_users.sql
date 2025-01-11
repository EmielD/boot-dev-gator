-- +goose Up
CREATE TABLE users(
    id UUID not null PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name varchar(255) not null UNIQUE
);

CREATE TABLE feeds(
    id UUID not null PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name varchar(255) not null,
    url varchar(255) not null UNIQUE,
    user_id UUID not null references users ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users;
DROP TABLE feeds;