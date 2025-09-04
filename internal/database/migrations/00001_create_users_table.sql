-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
    ID SERIAL NOT NULL UNIQUE,
    Username VARCHAR(127) NOT NULL UNIQUE,
    Password VARCHAR(127) NOT NULL,
    PRIMARY KEY (ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
