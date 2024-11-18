package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
	"github.com/shopspring/decimal"
)

type result struct {
	ID              int32           `csv:"resultId"`
	RaceID          int32           `csv:"raceId"`
	DriverID        int32           `csv:"driverId"`
	ConstructorID   int32           `csv:"constructorId"`
	Number          optionalNumber  `csv:"number"`
	Grid            int32           `csv:"grid"`
	Position        optionalNumber  `csv:"position"`
	PositionText    string          `csv:"positionText"`
	PositionOrder   int32           `csv:"positionOrder"`
	Points          decimal.Decimal `csv:"points"`
	Laps            int32           `csv:"laps"`
	Time            optionalString  `csv:"time"`
	Milliseconds    optionalNumber  `csv:"milliseconds"`
	FastestLap      optionalNumber  `csv:"fastestLap"`
	Rank            optionalNumber  `csv:"rank"`
	FastestLapTime  optionalString  `csv:"fastestLapTime"`
	FastestLapSpeed optionalString  `csv:"fastestLapSpeed"`
	StatusID        int32           `csv:"statusId"`
}

func (i Imager) loadResults(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/results.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error opening results CSV file: %w", err)
	}
	defer file.Close()

	var results []*result

	if err = gocsv.UnmarshalFile(file, &results); err != nil {
		return fmt.Errorf("Error marshaling results CSV file: %w", err)
	}

	records := []database.SaveResultsParams{}

	for _, result := range results {
		records = append(records, database.SaveResultsParams{
			ID:              result.ID,
			RaceID:          result.RaceID,
			DriverID:        result.DriverID,
			ConstructorID:   result.ConstructorID,
			Number:          result.Number.Value,
			Grid:            result.Grid,
			Position:        result.Position.Value,
			PosText:         result.PositionText,
			PosOrder:        result.PositionOrder,
			Points:          result.Points,
			Laps:            result.Laps,
			Time:            result.Time.Value,
			Milliseconds:    result.Milliseconds.Value,
			FastestLap:      result.FastestLap.Value,
			Rank:            result.Rank.Value,
			FastestLapTime:  result.FastestLapTime.Value,
			FastestLapSpeed: result.FastestLapSpeed.Value,
			StatusID:        &result.StatusID,
		})
	}

	_, err = i.db.WithTx(tx).SaveResults(ctx, records)
	if err != nil {
		return fmt.Errorf("Error saving results: %w", err)
	}

	fmt.Println("[Results] seeding complete")
	return nil
}

type sprintResult struct {
	ID              int32           `csv:"resultId"`
	RaceID          int32           `csv:"raceId"`
	DriverID        int32           `csv:"driverId"`
	ConstructorID   int32           `csv:"constructorId"`
	Number          int32           `csv:"number"`
	Grid            int32           `csv:"grid"`
	Position        optionalNumber  `csv:"position"`
	PositionText    string          `csv:"positionText"`
	PositionOrder   int32           `csv:"positionOrder"`
	Points          decimal.Decimal `csv:"points"`
	Laps            int32           `csv:"laps"`
	Time            optionalString  `csv:"time"`
	Milliseconds    optionalNumber  `csv:"milliseconds"`
	FastestLap      optionalNumber  `csv:"fastestLap"`
	Rank            optionalNumber  `csv:"rank"`
	FastestLapTime  optionalString  `csv:"fastestLapTime"`
	FastestLapSpeed optionalString  `csv:"fastestLapSpeed"`
	StatusID        int32           `csv:"statusId"`
}

func (i Imager) loadSprintResults(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/sprint_results.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error opening sprint results CSV file: %w", err)
	}
	defer file.Close()

	var results []*sprintResult

	if err = gocsv.UnmarshalFile(file, &results); err != nil {
		return fmt.Errorf("Error marshaling sprint results CSV file: %w", err)
	}

	records := []database.SaveSprintResultsParams{}

	for _, result := range results {
		records = append(records, database.SaveSprintResultsParams{
			ID:             result.ID,
			RaceID:         result.RaceID,
			DriverID:       result.DriverID,
			ConstructorID:  result.ConstructorID,
			Number:         result.Number,
			Grid:           result.Grid,
			Position:       result.Position.Value,
			PosText:        result.PositionText,
			PosOrder:       result.PositionOrder,
			Points:         result.Points,
			Laps:           result.Laps,
			Time:           result.Time.Value,
			Milliseconds:   result.Milliseconds.Value,
			FastestLap:     result.FastestLap.Value,
			FastestLapTime: result.FastestLapTime.Value,
			StatusID:       &result.StatusID,
		})
	}

	_, err = i.db.WithTx(tx).SaveSprintResults(ctx, records)
	if err != nil {
		return fmt.Errorf("Error saving sprint results: %w", err)
	}

	fmt.Println("[Sprint Results] seeding complete")
	return nil
}
