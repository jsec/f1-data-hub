package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type PitStopService struct {
	db *database.Queries
}

func NewPitStopService(db *database.Queries) PitStopService {
	return PitStopService{db}
}

func (s *PitStopService) SeedPitStops(ctx context.Context, tx pgx.Tx, records []database.SavePitStopsParams) error {
	_, err := s.db.WithTx(tx).SavePitStops(ctx, records)
	return err
}
