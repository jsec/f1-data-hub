package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsec/f1-data-hub/internal/database"
)

type Server struct {
	port int
	db   *database.Queries
}

func createServer(pool *pgxpool.Pool) *http.Server {
	// TODO: move this to an env variable
	port := 8080
	srv := &Server{
		port: port,
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
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("error getting database connection: %w", err)
	}
	defer pool.Close()

	server := createServer(pool)

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
