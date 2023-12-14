-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS concept (
    id CHAR(36) NOT NULL,
    player_id CHAR(36) NOT NULL,
    concept TEXT DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS concept;
-- +goose StatementEnd
