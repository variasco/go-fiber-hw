package database

import (
	"context"
	"variasco/go-fiber-hw/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)

	if err != nil {
		logger.Fatal().Err(err).Msg("Не удалось подключиться к базе данных")
		panic(err)
	}

	logger.Info().Msg("Соединение с базой данных установлено")
	return dbpool
}
