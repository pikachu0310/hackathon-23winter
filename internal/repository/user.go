package repository

import (
	"context"
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"

	"github.com/google/uuid"
)

type (
	// users table
	User struct {
		ID   uuid.UUID `db:"id"`
		Name string    `db:"name"`
	}

	CreateUserParams struct {
		Name string
	}

	CreateUserByIDParams struct {
		ID uuid.UUID
	}
)

func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {
	users := []*User{}
	if err := r.db.SelectContext(ctx, &users, "SELECT * FROM users"); err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, params CreateUserParams) (uuid.UUID, error) {
	userID := uuid.New()
	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id, name) VALUES (?, ?)", userID, params.Name); err != nil {
		return uuid.Nil, fmt.Errorf("insert user: %w", err)
	}

	return userID, nil
}

func (r *Repository) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	user := &User{}
	if err := r.db.GetContext(ctx, user, "SELECT * FROM users WHERE id = ?", userID); err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	return user, nil
}

func (r *Repository) CreateUserByUserID(ctx context.Context, params CreateUserByIDParams) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id, name) VALUES (?, ?)", params.ID, "test"); err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *Repository) ResetUsers() error {
	err := migration.ResetUserTable(r.db.DB)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ResetUsers() error {
	err := migration.ResetUserTable(r.db.DB)
	if err != nil {
		return err
	}

	return nil
}
