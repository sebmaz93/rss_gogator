-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds
DROP COLUMN updated_at;
-- +goose StatementEnd
