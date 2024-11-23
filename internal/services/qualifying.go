package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type QualifyingService struct {
	db *database.Queries
}

func NewQualifyingService(db *database.Queries) QualifyingService {
	return QualifyingService{db}
}

func (s *QualifyingService) SeedQualifyingResults(ctx context.Context, tx pgx.Tx, records []database.SaveQualifyingResultsParams) error {
	_, err := s.db.WithTx(tx).SaveQualifyingResults(ctx, records)
	return err
}
