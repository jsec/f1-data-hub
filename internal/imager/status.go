package imager

import (
	"context"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type status struct {
	ID     int    `csv:"statusId"`
	Status string `csv:"status"`
}

func (i Imager) loadStatuses(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/status.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening status CSV file: %w", err)
	}
	defer file.Close()

	var statuses []*status

	if err = gocsv.UnmarshalFile(file, &statuses); err != nil {
		return fmt.Errorf("error marshaling status CSV file: %w", err)
	}

	records := []database.SaveStatusesParams{}

	for _, status := range statuses {
		records = append(records, database.SaveStatusesParams{
			ID:     int32(status.ID),
			Status: status.Status,
		})
	}

	_, err = i.db.WithTx(tx).SaveStatuses(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving statuses: %w", err)
	}

	fmt.Println("[Statuses] seeding complete")
	return nil
}
