package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func NewPostgresPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
