package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type DriverService struct {
	db *database.Queries
}

func NewDriverService(db *database.Queries) DriverService {
	return DriverService{db}
}
