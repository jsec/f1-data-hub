package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type PitStopService struct {
	db *database.Queries
}

func NewPitStopService(db *database.Queries) PitStopService {
	return PitStopService{db}
}
