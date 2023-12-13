-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS battle (
    id CHAR(36) NOT NULL,
    my_kemono_id CHAR(36) NOT NULL,
    enemy_kemono_id CHAR(36) NOT NULL,
    text TEXT DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS battle;
-- +goose StatementEnd
