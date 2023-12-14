package repository

import (
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"
)

type (
	/*
		CREATE TABLE IF NOT EXISTS concept (
		    id CHAR(36) NOT NULL,
		    player_id CHAR(36) NOT NULL,
		    concept TEXT DEFAULT '',
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	*/

	Concept struct {
		ID        *uuid.UUID `db:"id" json:"id,omitempty"`
		PlayerID  *uuid.UUID `db:"player_id" json:"player_id,omitempty"`
		Concept   *string    `db:"concept" json:"concept,omitempty"`
		CreatedAt *string    `db:"created_at" json:"created_at,omitempty"`
	}
)

func (r *Repository) CreateConcept(playerId uuid.UUID, concept string) (conceptId uuid.UUID, err error) {
	conceptId = uuid.New()

	_, err = r.db.Exec(`
		INSERT INTO concept (id, player_id, concept)
		VALUES (?, ?, ?)
	`, conceptId, playerId, concept)
	if err != nil {
		return uuid.Nil, err
	}

	return conceptId, nil
}

func (r *Repository) GetConcept(conceptId uuid.UUID) (concept Concept, err error) {
	err = r.db.Get(&concept, `
		SELECT *
		FROM concept
		WHERE id = ?
	`, conceptId)
	if err != nil {
		return Concept{}, err
	}

	return concept, nil
}

func (r *Repository) GetConcepts() (concepts []Concept, err error) {
	err = r.db.Select(&concepts, `
		SELECT *
		FROM concept
	`)
	if err != nil {
		return nil, err
	}

	return concepts, nil
}

func (r *Repository) DeleteConcept(conceptId uuid.UUID) (err error) {
	_, err = r.db.Exec(`
		DELETE FROM concept
		WHERE id = ?
	`, conceptId)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetConceptsByPlayerId(playerId uuid.UUID) (concepts []Concept, err error) {
	err = r.db.Select(&concepts, `
		SELECT *
		FROM concept
		WHERE player_id = ?
	`, playerId)
	if err != nil {
		return nil, err
	}

	return concepts, nil
}

func (r *Repository) ResetConcepts() (err error) {
	err = migration.ResetConceptTable(r.db.DB)
	if err != nil {
		return err
	}

	return nil
}
