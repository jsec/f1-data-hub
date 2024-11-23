package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type ConstructorService struct {
	db *database.Queries
}

func NewConstructorService(db *database.Queries) ConstructorService {
	return ConstructorService{db}
}

func (s *ConstructorService) SeedConstructors(ctx context.Context, tx pgx.Tx, records []database.SaveConstructorsParams) error {
	_, err := s.db.WithTx(tx).SaveConstructors(ctx, records)
	return err
}

func (s *ConstructorService) SeedConstructorResults(ctx context.Context, tx pgx.Tx, records []database.SaveConstructorResultsParams) error {
	_, err := s.db.WithTx(tx).SaveConstructorResults(ctx, records)
	return err
}

func (s *ConstructorService) SeedConstructorStandings(ctx context.Context, tx pgx.Tx, records []database.SaveConstructorStandingsParams) error {
	_, err := s.db.WithTx(tx).SaveConstructorStandings(ctx, records)
	return err
}
