package database

import (
	"context"
	"fmt"

	"github.com/cdlinkin/system-booking/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(cfg *config.DBConfig) (*pgxpool.Pool, error) {
	connDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name)

	pool, err := pgxpool.New(context.Background(), connDB)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать connection pool %w", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к БД: %w", err)
	}

	return pool, nil
}
