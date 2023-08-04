-- +goose Up
-- +goose StatementBegin
ALTER TABLE members ADD COLUMN mail_send timestamptz;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE members DROP COLUMN mail_send;
-- +goose StatementEnd
