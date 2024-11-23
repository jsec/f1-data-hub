package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type DriverService struct {
	db *database.Queries
}

func NewDriverService(db *database.Queries) DriverService {
	return DriverService{db}
}

func (s *DriverService) SeedDrivers(ctx context.Context, tx pgx.Tx, records []database.SaveDriversParams) error {
	_, err := s.db.WithTx(tx).SaveDrivers(ctx, records)
	return err
}

func (s *DriverService) SeedDriverStandings(ctx context.Context, tx pgx.Tx, records []database.SaveDriverStandingsParams) error {
	_, err := s.db.WithTx(tx).SaveDriverStandings(ctx, records)
	return err
}
