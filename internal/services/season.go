package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
	"github.com/shopspring/decimal"
)

type SeasonService struct {
	db *database.Queries
}

func NewSeasonService(db *database.Queries) SeasonService {
	return SeasonService{db}
}

type DriverStanding struct {
	ID        int32           `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Points    decimal.Decimal `json:"points"`
}

func (s *SeasonService) SeedSeasons(ctx context.Context, tx pgx.Tx, records []database.SaveSeasonsParams) error {
	_, err := s.db.WithTx(tx).SaveSeasons(ctx, records)
	return err
}

func (s *SeasonService) GetDriverStandingsByYear(ctx context.Context, year int32) ([]DriverStanding, error) {
	result, err := s.db.GetDriverStandingsByYear(ctx, year)
	if err != nil {
		return nil, fmt.Errorf("error retrieving driver standings: %w", err)
	}

	var standings []DriverStanding

	for _, row := range result {
		standings = append(standings, DriverStanding{
			ID:        row.DriverID,
			FirstName: row.FirstName,
			LastName:  row.LastName,
			Points:    row.Points,
		})
	}

	return standings, nil
}
