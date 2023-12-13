package repository

import (
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"
)

type (
	/*
		CREATE TABLE IF NOT EXISTS battle (
		    id VARCHAR(36) NOT NULL,
		    my_kemono_id VARCHAR(36) NOT NULL,
		    enemy_kemono_id VARCHAR(36) NOT NULL,
		    text TEXT DEFAULT '',
		    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	*/

	Battle struct {
		ID            *uuid.UUID `db:"id" json:"id,omitempty"`
		MyKemonoID    *uuid.UUID `db:"my_kemono_id" json:"my_kemono_id,omitempty"`
		EnemyKemonoID *uuid.UUID `db:"enemy_kemono_id" json:"enemy_kemono_id,omitempty"`
		Text          *string    `db:"text" json:"text,omitempty"`
		CreatedAt     *string    `db:"created_at" json:"created_at,omitempty"`
	}
)

func (r *Repository) CreateBattle(MyKemonoID uuid.UUID, EnemyKemonoID uuid.UUID, text string) (battleId uuid.UUID, err error) {
	battleId = uuid.New()

	_, err = r.db.Exec(`
		INSERT INTO battle (id, my_kemono_id, enemy_kemono_id, text)
		VALUES (?, ?, ?, ?)
	`, battleId, MyKemonoID, EnemyKemonoID, text)
	if err != nil {
		return uuid.Nil, err
	}

	return battleId, nil
}

func (r *Repository) GetBattle(battleId uuid.UUID) (battle Battle, err error) {
	err = r.db.Get(&battle, `
		SELECT *
		FROM battle
		WHERE id = ?
	`, battleId)
	if err != nil {
		return Battle{}, err
	}

	return battle, nil
}

func (r *Repository) GetBattles() (battles []Battle, err error) {
	err = r.db.Select(&battles, `
		SELECT *
		FROM battle
	`)
	if err != nil {
		return nil, err
	}

	return battles, nil
}

func (r *Repository) UpdateBattleText(battleId uuid.UUID, text string) (err error) {
	_, err = r.db.Exec(`
		UPDATE battle
		SET text = ?
		WHERE id = ?
	`, text, battleId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ResetBattles() error {
	err := migration.ResetBattleTable(r.db.DB)
	if err != nil {
		return err
	}

	return nil
}
