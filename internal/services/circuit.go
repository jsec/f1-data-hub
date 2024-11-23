package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type CircuitService struct {
	db *database.Queries
}

func NewCircuitService(db *database.Queries) CircuitService {
	return CircuitService{db}
}

func (c *CircuitService) SeedCircuits(ctx context.Context, tx pgx.Tx, records []database.SaveCircuitsParams) error {
	_, err := c.db.SaveCircuits(ctx, records)
	return err
}
