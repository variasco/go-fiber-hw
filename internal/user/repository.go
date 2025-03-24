package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UserRepository struct {
	dbpool *pgxpool.Pool
	logger *zerolog.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger) *UserRepository {
	return &UserRepository{
		dbpool: dbpool,
		logger: logger,
	}
}

func (repo *UserRepository) create(form RegisterForm) (*User, error) {
	query := `INSERT INTO users (name, email, password, created_at) 
	          VALUES (@name, @email, @password, @created_at) 
	          RETURNING id, name, email, password, created_at;`

	hashedPassword := form.Password
	args := pgx.NamedArgs{
		"name":       form.Name,
		"email":      form.Email,
		"password":   hashedPassword,
		"created_at": time.Now().UTC(),
	}

	var user User
	err := repo.dbpool.QueryRow(context.Background(), query, args).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	return &user, nil
}
