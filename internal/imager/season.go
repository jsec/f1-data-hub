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

func (i Imager) loadSeasons(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/seasons.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening seasons CSV file: %w", err)
	}
	defer file.Close()

	var seasons []*season

	if err = gocsv.UnmarshalFile(file, &seasons); err != nil {
		return fmt.Errorf("error marshaling seasons CSV file: %w", err)
	}

	records := []database.SaveSeasonsParams{}

	for _, season := range seasons {
		records = append(records, database.SaveSeasonsParams{
			Year: int32(season.Year),
			Url:  season.Url,
		})
	}

	_, err = i.db.WithTx(tx).SaveSeasons(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving seasons: %w", err)
	}

	fmt.Println("[Seasons] seeding complete")
	return nil
}
