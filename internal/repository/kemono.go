package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type (
	/*
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
	*/

	Kemono struct {
		ID            uuid.UUID `db:"id"`
		Image         []byte    `db:"image"`
		Prompt        string    `db:"prompt"`
		Name          string    `db:"name"`
		Description   string    `db:"description"`
		CharacterChip int       `db:"character_chip"`
		IsPlayer      bool      `db:"is_player"`
		PlayerID      uuid.UUID `db:"player_id"`
		IsOwned       bool      `db:"is_owned"`
		OwnerID       uuid.UUID `db:"owner_id"`
		IsInField     bool      `db:"is_in_field"`
		IsBoss        bool      `db:"is_boss"`
		Field         int       `db:"field"`
		X             int       `db:"x"`
		Y             int       `db:"y"`
		HasParent     bool      `db:"has_parent"`
		Parent1ID     uuid.UUID `db:"parent1_id"`
		Parent2ID     uuid.UUID `db:"parent2_id"`
		HasChild      bool      `db:"has_child"`
		ChildID       uuid.UUID `db:"child_id"`
		Hp            int       `db:"hp"`
		Attack        int       `db:"attack"`
		Defense       int       `db:"defense"`
		CreatedAt     string    `db:"created_at"`
	}
)

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
	if kemono.ID == uuid.Nil {
		kemono.ID = uuid.New() // IDを設定
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

	return kemono.ID, nil
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
	if err := r.db.SelectContext(ctx, &kemono, "SELECT * FROM kemono WHERE field = ?", field); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return kemono, nil
}
