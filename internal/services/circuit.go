package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type CircuitService struct {
	db *database.Queries
}

func NewCircuitService(db *database.Queries) CircuitService {
	return CircuitService{db}
}
