package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"
	"reflect"
	"strings"
)

type (
	/*
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
	*/

	ConceptsText string
	Concepts     []string

	Kemono struct {
		ID            *uuid.UUID    `db:"id"`
		Image         []byte        `db:"image"`
		Prompt        *string       `db:"prompt"`
		Concepts      *ConceptsText `db:"concepts"`
		Name          *string       `db:"name"`
		Description   *string       `db:"description"`
		CharacterChip *int          `db:"character_chip"`
		IsPlayer      *bool         `db:"is_player"`
		PlayerID      *uuid.UUID    `db:"player_id"`
		IsOwned       *bool         `db:"is_owned"`
		OwnerID       *uuid.UUID    `db:"owner_id"`
		IsInField     *bool         `db:"is_in_field"`
		IsBoss        *bool         `db:"is_boss"`
		Field         *int          `db:"field"`
		X             *int          `db:"x"`
		Y             *int          `db:"y"`
		HasParent     *bool         `db:"has_parent"`
		Parent1ID     *uuid.UUID    `db:"parent1_id"`
		Parent2ID     *uuid.UUID    `db:"parent2_id"`
		HasChild      *bool         `db:"has_child"`
		ChildID       *uuid.UUID    `db:"child_id"`
		MaxHp         *int          `db:"max_hp"`
		Hp            *int          `db:"hp"`
		Attack        *int          `db:"attack"`
		Defense       *int          `db:"defense"`
		CreatedAt     *string       `db:"created_at"`
	}

	KemonoParams struct {
		ID            uuid.UUID
		Image         []byte
		Prompt        string
		Concepts      ConceptsText
		Name          string
		Description   string
		CharacterChip int
		IsPlayer      bool
		PlayerID      uuid.UUID
		IsOwned       bool
		OwnerID       uuid.UUID
		IsInField     bool
		IsBoss        bool
		Field         int
		X             int
		Y             int
		HasParent     bool
		Parent1ID     uuid.UUID
		Parent2ID     uuid.UUID
		HasChild      bool
		ChildID       uuid.UUID
		MaxHp         int
		Hp            int
		Attack        int
		Defense       int
		CreatedAt     string
	}
)

func (c ConceptsText) Concepts() Concepts {
	// "a,b,c" -> ["a", "b", "c"]
	return strings.Split(c.String(), ",")
}

func (c ConceptsText) String() string {
	return string(c)
}

func (c Concepts) String() string {
	// ["a", "b", "c"] -> "a,b,c"
	return strings.Join(c, ",")
}

func (kemonoParams *KemonoParams) ToKemono() *Kemono {
	return &Kemono{
		ID:            &kemonoParams.ID,
		Image:         kemonoParams.Image,
		Prompt:        &kemonoParams.Prompt,
		Concepts:      &kemonoParams.Concepts,
		Name:          &kemonoParams.Name,
		Description:   &kemonoParams.Description,
		CharacterChip: &kemonoParams.CharacterChip,
		IsPlayer:      &kemonoParams.IsPlayer,
		PlayerID:      &kemonoParams.PlayerID,
		IsOwned:       &kemonoParams.IsOwned,
		OwnerID:       &kemonoParams.OwnerID,
		IsInField:     &kemonoParams.IsInField,
		IsBoss:        &kemonoParams.IsBoss,
		Field:         &kemonoParams.Field,
		X:             &kemonoParams.X,
		Y:             &kemonoParams.Y,
		HasParent:     &kemonoParams.HasParent,
		Parent1ID:     &kemonoParams.Parent1ID,
		Parent2ID:     &kemonoParams.Parent2ID,
		HasChild:      &kemonoParams.HasChild,
		ChildID:       &kemonoParams.ChildID,
		MaxHp:         &kemonoParams.MaxHp,
		Hp:            &kemonoParams.Hp,
		Attack:        &kemonoParams.Attack,
		Defense:       &kemonoParams.Defense,
		CreatedAt:     &kemonoParams.CreatedAt,
	}
}

func isZeroValue(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func addFieldsForKemono(queryBase *string, values *string, args *[]interface{}, kemono *Kemono) {
	t := reflect.TypeOf(*kemono)
	v := reflect.ValueOf(*kemono)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		if !isZeroValue(value) || (field.Name == "Image" && len(kemono.Image) > 0) {
			dbTag := field.Tag.Get("db")
			if dbTag != "" && dbTag != "id" && dbTag != "created_at" {
				*queryBase += fmt.Sprintf(", %s", dbTag)
				if values != nil {
					*values += ", ?"
				} else {
					*queryBase += " = ?"
				}
				*args = append(*args, value)
			}
		}
	}
}

func (r *Repository) CreateKemono(ctx context.Context, kemono *Kemono) (uuid.UUID, error) {
	if kemono.ID == nil {
		newUUID := uuid.New()
		kemono.ID = &newUUID // IDを設定
	}
	query := "INSERT INTO kemono (id"
	values := "(?"
	args := []interface{}{kemono.ID}

	addFieldsForKemono(&query, &values, &args, kemono)

	query += ") VALUES " + values + ")"
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert kemono: %w", err)
	}

	return *kemono.ID, nil
}

func (r *Repository) UpdateKemono(ctx context.Context, kemono *Kemono) error {
	query := "UPDATE kemono SET "
	args := []interface{}{}

	addFieldsForKemono(&query, nil, &args, kemono)

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id = ?"
	args = append(args, kemono.ID)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("update kemono: %w", err)
	}
	return nil
}

func (r *Repository) GetKemonos(ctx context.Context) ([]Kemono, error) {
	var kemonos []Kemono
	if err := r.db.SelectContext(ctx, &kemonos, "SELECT * FROM kemono"); err != nil {
		return nil, fmt.Errorf("select kemonos: %w", err)
	}

	return kemonos, nil
}

func (r *Repository) GetKemono(ctx context.Context, kemonoID uuid.UUID) (*Kemono, error) {
	var kemono Kemono
	if err := r.db.GetContext(ctx, &kemono, "SELECT * FROM kemono WHERE id = ?", kemonoID); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return &kemono, nil
}

func (r *Repository) GetKemonosByField(ctx context.Context, field int) ([]Kemono, error) {
	var kemono []Kemono
	if err := r.db.SelectContext(ctx, &kemono, "SELECT * FROM kemono WHERE field = ? AND is_in_field = TRUE", field); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return kemono, nil
}

func (r *Repository) GetKemonoByOwnerId(ctx context.Context, ownerID uuid.UUID) ([]Kemono, error) {
	var kemono []Kemono
	if err := r.db.SelectContext(ctx, &kemono, "SELECT * FROM kemono WHERE owner_id = ?", ownerID); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return kemono, nil
}

func (r *Repository) ResetKemonos() error {
	err := migration.ResetKemonoTable(r.db.DB)
	if err != nil {
		return err
	}

	return nil
}
