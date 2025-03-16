package services

import (
	"github.com/jsec/f1-data-hub/internal/database"
)

type QualifyingService struct {
	db *database.Queries
}

func NewQualifyingService(db *database.Queries) QualifyingService {
	return QualifyingService{db}
}
