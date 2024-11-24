package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsec/f1-data-hub/internal/config"
	"github.com/jsec/f1-data-hub/internal/database"
)

type Server struct {
	port int
	db   *database.Queries
}

func createServer(pool *pgxpool.Pool, cfg config.Config) *http.Server {
	srv := &Server{
		port: cfg.Port,
		db:   database.New(pool),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.port),
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func Run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("error parsing configuration: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("error getting database connection: %w", err)
	}
	defer pool.Close()

	server := createServer(pool, cfg)

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
