package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"
	"reflect"
	"strings"
)

func isZeroValue(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func addFieldsForKemono(queryBase *string, values *string, args *[]interface{}, kemono *domains.Kemono) {
	t := reflect.TypeOf(*kemono)
	v := reflect.ValueOf(*kemono)
	first := true

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		if !isZeroValue(value) || (field.Name == "Image" && len(kemono.Image) > 0) {
			dbTag := field.Tag.Get("db")
			if dbTag != "" && dbTag != "id" && dbTag != "created_at" {
				if values != nil {
					*queryBase += fmt.Sprintf(", %s", dbTag)
					*values += ", ?"
				} else {
					if first {
						*queryBase += fmt.Sprintf("%s", dbTag)
					} else {
						*queryBase += fmt.Sprintf(", %s", dbTag)
					}
					*queryBase += " = ?"
				}
				*args = append(*args, value)
				first = false
			}
		}
	}
}

func (r *Repository) CreateKemono(ctx context.Context, kemono *domains.Kemono) (uuid.UUID, error) {
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

func (r *Repository) UpdateKemono(ctx context.Context, kemono *domains.Kemono) error {
	query := "UPDATE kemono SET "
	var args []interface{}

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

func (r *Repository) GetKemonos(ctx context.Context) ([]domains.Kemono, error) {
	var kemonos []domains.Kemono
	if err := r.db.SelectContext(ctx, &kemonos, "SELECT * FROM kemono"); err != nil {
		return nil, fmt.Errorf("select kemonos: %w", err)
	}

	return kemonos, nil
}

func (r *Repository) GetKemono(ctx context.Context, kemonoID uuid.UUID) (*domains.Kemono, error) {
	var kemono domains.Kemono
	if err := r.db.GetContext(ctx, &kemono, "SELECT * FROM kemono WHERE id = ?", kemonoID); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return &kemono, nil
}

func (r *Repository) GetKemonosByField(ctx context.Context, field int) ([]domains.Kemono, error) {
	var kemono []domains.Kemono
	if err := r.db.SelectContext(ctx, &kemono, "SELECT * FROM kemono WHERE field = ? AND is_in_field = TRUE", field); err != nil {
		return nil, fmt.Errorf("select kemono: %w", err)
	}

	return kemono, nil
}

func (r *Repository) GetKemonoByOwnerId(ctx context.Context, ownerID uuid.UUID) ([]domains.Kemono, error) {
	var kemono []domains.Kemono
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
