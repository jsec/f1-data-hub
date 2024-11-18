package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type qualifying struct {
	ID            int32          `csv:"qualifyId"`
	RaceID        int32          `csv:"raceId"`
	DriverID      int32          `csv:"driverId"`
	ConstructorID int32          `csv:"constructorId"`
	Number        int32          `csv:"number"`
	Position      int32          `csv:"position"`
	Q1            optionalString `csv:"q1"`
	Q2            optionalString `csv:"q2"`
	Q3            optionalString `csv:"q3"`
}

func (i Imager) loadQualifying(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/qualifying.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error opening qualifying CSV file: %w", err)
	}
	defer file.Close()

	var results []*qualifying

	if err = gocsv.UnmarshalFile(file, &results); err != nil {
		return fmt.Errorf("Error marshaling qualifying CSV file: %w", err)
	}

	records := []database.SaveQualifyingResultsParams{}

	for _, result := range results {
		records = append(records, database.SaveQualifyingResultsParams{
			ID:            result.ID,
			RaceID:        result.RaceID,
			DriverID:      result.DriverID,
			ConstructorID: result.ConstructorID,
			Number:        result.Number,
			Position:      &result.Position,
			Q1:            result.Q1.Value,
			Q2:            result.Q2.Value,
			Q3:            result.Q3.Value,
		})
	}

	_, err = i.db.WithTx(tx).SaveQualifyingResults(ctx, records)
	if err != nil {
		return fmt.Errorf("Error saving qualifying results: %w", err)
	}

	fmt.Println("[Qualifying] seeding complete")
	return nil
}
