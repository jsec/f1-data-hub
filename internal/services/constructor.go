package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type ConstructorService struct {
	db *database.Queries
}

func NewConstructorService(db *database.Queries) ConstructorService {
	return ConstructorService{db}
}
