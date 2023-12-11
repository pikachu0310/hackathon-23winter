-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(36) NOT NULL,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS kemono (
    id VARCHAR(36) NOT NULL,
    image MEDIUMBLOB NOT NULL,
    prompt VARCHAR(255),
    name VARCHAR(255),
    description VARCHAR(255),
    character_chip INT,
    is_player BOOLEAN NOT NULL DEFAULT FALSE,
    player_id VARCHAR(36),
    is_owned BOOLEAN NOT NULL DEFAULT FALSE,
    owner_id VARCHAR(36),
    is_in_field BOOLEAN NOT NULL DEFAULT TRUE,
    is_boss BOOLEAN NOT NULL DEFAULT FALSE,
    field INT,
    x INT,
    y INT,
    has_parent BOOLEAN NOT NULL DEFAULT FALSE,
    parent1_id VARCHAR(36),
    parent2_id VARCHAR(36),
    has_child BOOLEAN NOT NULL DEFAULT FALSE,
    child_id VARCHAR(36),
    hp INT,
    attack INT,
    defense INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
