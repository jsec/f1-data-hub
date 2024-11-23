package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type LapTimeService struct {
	db *database.Queries
}

func NewLapTimeService(db *database.Queries) LapTimeService {
	return LapTimeService{db}
}

func (s *LapTimeService) SeedLapTimes(ctx context.Context, tx pgx.Tx, records []database.SaveLapTimesParams) error {
	_, err := s.db.WithTx(tx).SaveLapTimes(ctx, records)
	return err
}
