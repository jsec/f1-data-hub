package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type RaceService struct {
	db *database.Queries
}

func NewRaceService(db *database.Queries) RaceService {
	return RaceService{db}
}

func (s *RaceService) SeedRaces(ctx context.Context, tx pgx.Tx, records []database.SaveRacesParams) error {
	_, err := s.db.WithTx(tx).SaveRaces(ctx, records)
	return err
}
