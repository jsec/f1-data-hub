package repository

import (
	"context"
	"fmt"

	"github.com/jsec/f1-data-hub/internal/database"
	"github.com/shopspring/decimal"
)

func NewSeasonRepository(db *database.Queries) Repository {
	return Repository{db: db}
}

type DriverStanding struct {
	ID        int32           `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Points    decimal.Decimal `json:"points"`
}

func (r *Repository) GetDriverStandingsByYear(ctx context.Context, year int32) ([]DriverStanding, error) {
	result, err := r.db.GetDriverStandingsByYear(ctx, year)
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
