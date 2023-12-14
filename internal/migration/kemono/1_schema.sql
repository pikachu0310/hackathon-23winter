-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS kemono (
    id CHAR(36) NOT NULL,
    image MEDIUMBLOB NOT NULL,
    prompt TEXT DEFAULT '',
    concepts TEXT DEFAULT '',
    name TEXT DEFAULT '',
    description TEXT DEFAULT '',
    character_chip INT DEFAULT -1,
    is_player BOOLEAN NOT NULL DEFAULT FALSE DEFAULT FALSE,
    player_id CHAR(36) DEFAULT '',
    is_owned BOOLEAN NOT NULL DEFAULT FALSE DEFAULT FALSE,
    owner_id CHAR(36) DEFAULT '',
    is_in_field BOOLEAN NOT NULL DEFAULT TRUE DEFAULT TRUE,
    is_boss BOOLEAN NOT NULL DEFAULT FALSE,
    field INT DEFAULT -1,
    x INT DEFAULT -1,
    y INT DEFAULT -1,
    has_parent BOOLEAN NOT NULL DEFAULT FALSE,
    parent1_id CHAR(36) DEFAULT '',
    parent2_id CHAR(36) DEFAULT '',
    has_child BOOLEAN NOT NULL DEFAULT FALSE,
    child_id CHAR(36) DEFAULT '',
    max_hp INT DEFAULT -1,
    hp INT DEFAULT -1,
    attack INT DEFAULT -1,
    defense INT DEFAULT -1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS kemono;
-- +goose StatementEnd
