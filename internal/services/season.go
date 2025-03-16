package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type SeasonService struct {
	db *database.Queries
}

func NewSeasonService(db *database.Queries) SeasonService {
	return SeasonService{db}
}
