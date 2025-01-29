-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS emails
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER      NOT NULL,
    email   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    sent_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS emails;
-- +goose StatementEnd
