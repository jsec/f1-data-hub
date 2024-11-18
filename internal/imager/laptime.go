package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type lapTime struct {
	RaceID       int32  `csv:"raceId"`
	DriverID     int32  `csv:"driverId"`
	Lap          int32  `csv:"lap"`
	Position     int32  `csv:"position"`
	Time         string `csv:"time"`
	Milliseconds int32  `csv:"milliseconds"`
}

func (i Imager) loadLapTimes(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/lap_times.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening lap times CSV file: %w", err)
	}
	defer file.Close()

	var lapTimes []*lapTime

	if err = gocsv.UnmarshalFile(file, &lapTimes); err != nil {
		return fmt.Errorf("error marshaling lap times CSV file: %w", err)
	}

	records := []database.SaveLapTimesParams{}

	for _, lapTime := range lapTimes {
		records = append(records, database.SaveLapTimesParams{
			RaceID:       lapTime.RaceID,
			DriverID:     lapTime.DriverID,
			Lap:          lapTime.Lap,
			Position:     &lapTime.Position,
			Time:         &lapTime.Time,
			Milliseconds: &lapTime.Milliseconds,
		})
	}

	_, err = i.db.WithTx(tx).SaveLapTimes(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving lap times: %w", err)
	}

	fmt.Println("[Lap Times] seeding complete")
	return nil
}
