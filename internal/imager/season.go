package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type season struct {
	Year int    `csv:"year"`
	Url  string `csv:"url"`
}

func (i imager) loadSeasons(ctx context.Context, tx pgx.Tx) error {
	spinner := i.spinners.AddSpinner("Seeding seasons")

	file, err := os.OpenFile("data/seasons.csv", os.O_RDONLY, 0600)
	if err != nil {
		spinner.Error()
		return fmt.Errorf("error opening seasons CSV file: %w", err)
	}
	defer file.Close()

	var seasons []*season

	if err = gocsv.UnmarshalFile(file, &seasons); err != nil {
		spinner.Error()
		return fmt.Errorf("error marshaling seasons CSV file: %w", err)
	}

	records := []database.SaveSeasonsParams{}

	for _, season := range seasons {
		records = append(records, database.SaveSeasonsParams{
			Year: int32(season.Year),
			Url:  season.Url,
		})
	}

	if err = i.seasonService.SeedSeasons(ctx, tx, records); err != nil {
		spinner.Error()
		return fmt.Errorf("error saving seasons: %w", err)
	}

	spinner.Complete()
	return nil
}
