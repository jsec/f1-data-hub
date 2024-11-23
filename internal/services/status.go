package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type StatusService struct {
	db *database.Queries
}

func NewStatusService(db *database.Queries) StatusService {
	return StatusService{db}
}

func (s *StatusService) SeedStatuses(ctx context.Context, tx pgx.Tx, records []database.SaveStatusesParams) error {
	_, err := s.db.WithTx(tx).SaveStatuses(ctx, records)
	return err
}
