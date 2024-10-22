package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jsec/f1-data-hub/internal/database"
)

type Server struct {
	port int
	db   *database.Queries
}

func NewServer(pool *pgxpool.Pool) *http.Server {
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
