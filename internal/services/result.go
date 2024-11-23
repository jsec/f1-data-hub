package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type ResultsService struct {
	db *database.Queries
}

func NewResultsService(db *database.Queries) ResultsService {
	return ResultsService{db}
}

func (s *ResultsService) SeedResults(ctx context.Context, tx pgx.Tx, records []database.SaveResultsParams) error {
	_, err := s.db.WithTx(tx).SaveResults(ctx, records)
	return err
}

func (s *ResultsService) SeedSprintResults(ctx context.Context, tx pgx.Tx, records []database.SaveSprintResultsParams) error {
	_, err := s.db.WithTx(tx).SaveSprintResults(ctx, records)
	return err
}
