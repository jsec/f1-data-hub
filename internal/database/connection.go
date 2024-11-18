package database

import (
	"context"
	"fmt"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	connString := os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %w", err)
	}

	config.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxdecimal.Register(c.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating connection: %w", err)
	}

	return pool, nil
}
