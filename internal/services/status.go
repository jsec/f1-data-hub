package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type StatusService struct {
	db *database.Queries
}

func NewStatusService(db *database.Queries) StatusService {
	return StatusService{db}
}
