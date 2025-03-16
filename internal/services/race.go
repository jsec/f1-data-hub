package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type RaceService struct {
	db *database.Queries
}

func NewRaceService(db *database.Queries) RaceService {
	return RaceService{db}
}
