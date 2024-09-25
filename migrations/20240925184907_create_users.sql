-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id bigserial NOT NULL,
    email varchar(255) NOT NULL unique,
    encrypted_password varchar(255) NOT NULL
);
ALTER TABLE users ADD PRIMARY KEY (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
