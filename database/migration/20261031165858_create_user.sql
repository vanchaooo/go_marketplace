-- +goose Up
CREATE TABLE users
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(200) not null check(char_length(name) > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE users
