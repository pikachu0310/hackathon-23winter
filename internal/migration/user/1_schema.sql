-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id CHAR(36) NOT NULL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS kemono;
-- +goose StatementEnd
