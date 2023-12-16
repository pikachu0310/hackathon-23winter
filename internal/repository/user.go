package repository

import (
	"context"
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/migration"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type (
	// users table
	User struct {
		ID        uuid.UUID `db:"id"`
		Name      string    `db:"name"`
		CreatedAt string    `db:"created_at"`
	}

	CreateUserParams struct {
		Name     string
		Password string
	}

	CreateUserByIDParams struct {
		ID uuid.UUID
	}
)

func (r *Repository) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	if err := r.db.SelectContext(ctx, &users, "SELECT id, name, created_at FROM users"); err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, params CreateUserParams) (uuid.UUID, error) {
	userID := uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("hash password: %w", err)
	}

	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id, name, password) VALUES (?, ?, ?)", userID, params.Name, hashedPassword); err != nil {
		return uuid.Nil, fmt.Errorf("insert user: %w", err)
	}

	return userID, nil
}

func (r *Repository) GetHashedPassword(ctx context.Context, userName string) ([]byte, error) {
	var storedPassword string

	err := r.db.QueryRowContext(ctx, "SELECT password FROM users WHERE name = ?", userName).Scan(&storedPassword)
	if err != nil {
		return nil, fmt.Errorf("select password: %w", err)
	}

	return []byte(storedPassword), nil
}

func (r *Repository) GetUserID(ctx context.Context, userName string) (uuid.UUID, error) {
	var userID uuid.UUID

	err := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE name = ?", userName).Scan(&userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("select id: %w", err)
	}

	return userID, nil
}

func (r *Repository) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	var user User
	if err := r.db.GetContext(ctx, &user, "SELECT id, name, created_at FROM users WHERE id = ?", userID); err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	return &user, nil
}

func (r *Repository) CreateUserByUserID(ctx context.Context, params CreateUserByIDParams) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO users (id, name, password) VALUES (?, ?, ?)", params.ID, params.ID, "test"); err != nil {
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
