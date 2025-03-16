package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type ResultsService struct {
	db *database.Queries
}

func NewResultsService(db *database.Queries) ResultsService {
	return ResultsService{db}
}
