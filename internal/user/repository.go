package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
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

func (repo *UserRepository) create(form RegisterForm) error {
	query := `INSERT INTO users (name, email, password, created_at) 
	          VALUES (@name, @email, @password, @created_at);`
	hashedPassword, err := hashPassword(form.Password)
	if err != nil {
		return fmt.Errorf("не удалось захэшировать пароль:%w", err)
	}
	args := pgx.NamedArgs{
		"name":       form.Name,
		"email":      form.Email,
		"password":   hashedPassword,
		"created_at": time.Now().UTC(),
	}
	_, err = repo.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("не удалось создать пользователя, %w", err)
	}

	return nil
}

func (repo *UserRepository) getByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = @email;`
	args := pgx.NamedArgs{
		"email": email,
	}

	var user User
	err := repo.dbpool.QueryRow(context.Background(), query, args).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	return &user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
