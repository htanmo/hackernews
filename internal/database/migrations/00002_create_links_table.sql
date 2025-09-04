-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Links (
    ID SERIAL NOT NULL UNIQUE,
    Title VARCHAR(255),
    Address VARCHAR(255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Links;
-- +goose StatementEnd
