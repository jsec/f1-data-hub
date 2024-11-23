package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type race struct {
	ID         int32            `csv:"raceId"`
	Year       int32            `csv:"year"`
	Round      int32            `csv:"round"`
	CircuitID  int32            `csv:"circuitId"`
	Name       string           `csv:"name"`
	Date       dateOnly         `csv:"date"`
	Time       optionalTimeOnly `csv:"time"`
	URL        string           `csv:"url"`
	Fp1Date    optionalDateOnly `csv:"fp1_date"`
	Fp1Time    optionalTimeOnly `csv:"fp1_time"`
	Fp2Date    optionalDateOnly `csv:"fp2_date"`
	Fp2Time    optionalTimeOnly `csv:"fp2_time"`
	Fp3Date    optionalDateOnly `csv:"fp3_date"`
	Fp3Time    optionalTimeOnly `csv:"fp3_time"`
	QualiDate  optionalDateOnly `csv:"quali_date"`
	QualiTime  optionalTimeOnly `csv:"quali_time"`
	SprintDate optionalDateOnly `csv:"sprint_date"`
	SprintTime optionalTimeOnly `csv:"sprint_time"`
}

func (i Imager) loadRaces(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/races.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening races CSV file: %w", err)
	}
	defer file.Close()

	var races []*race

	if err = gocsv.UnmarshalFile(file, &races); err != nil {
		return fmt.Errorf("error marshaling races CSV file: %w", err)
	}

	records := []database.SaveRacesParams{}

	for _, race := range races {
		records = append(records, database.SaveRacesParams{
			ID:         race.ID,
			Year:       race.Year,
			Round:      race.Round,
			CircuitID:  race.CircuitID,
			Name:       race.Name,
			Date:       race.Date.Value,
			Time:       race.Time.Value,
			Url:        &race.URL,
			Fp1Date:    race.Fp1Date.Value,
			Fp1Time:    race.Fp1Time.Value,
			Fp2Date:    race.Fp2Date.Value,
			Fp2Time:    race.Fp2Time.Value,
			Fp3Date:    race.Fp3Date.Value,
			Fp3Time:    race.Fp3Time.Value,
			QualiDate:  race.QualiDate.Value,
			QualiTime:  race.QualiTime.Value,
			SprintDate: race.SprintDate.Value,
			SprintTime: race.SprintTime.Value,
		})
	}

	if err = i.raceService.SeedRaces(ctx, tx, records); err != nil {
		return fmt.Errorf("error saving races: %w", err)
	}

	fmt.Println("[Races] seeding complete")
	return nil
}
