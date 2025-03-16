package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type LapTimeService struct {
	db *database.Queries
}

func NewLapTimeService(db *database.Queries) LapTimeService {
	return LapTimeService{db}
}
