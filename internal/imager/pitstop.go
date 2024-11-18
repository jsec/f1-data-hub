package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type pitStop struct {
	RaceID       int32    `csv:"raceId"`
	DriverID     int32    `csv:"driverId"`
	Stop         int32    `csv:"stop"`
	Lap          int32    `csv:"lap"`
	Time         timeOnly `csv:"time"`
	Duration     string   `csv:"duration"`
	Milliseconds int32    `csv:"milliseconds"`
}

func (i Imager) loadPitStops(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/pit_stops.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening pit stops CSV file: %w", err)
	}
	defer file.Close()

	var pitStops []*pitStop

	if err = gocsv.UnmarshalFile(file, &pitStops); err != nil {
		return fmt.Errorf("error marshaling pit stops CSV file: %w", err)
	}

	records := []database.SavePitStopsParams{}

	for _, pitStop := range pitStops {
		records = append(records, database.SavePitStopsParams{
			RaceID:       pitStop.RaceID,
			DriverID:     pitStop.DriverID,
			Stop:         pitStop.Stop,
			Lap:          pitStop.Lap,
			Time:         pitStop.Time.Value,
			Duration:     &pitStop.Duration,
			Milliseconds: &pitStop.Milliseconds,
		})
	}

	_, err = i.db.WithTx(tx).SavePitStops(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving pit stops: %w", err)
	}

	fmt.Println("[Pit Stops] seeding complete")
	return nil
}
