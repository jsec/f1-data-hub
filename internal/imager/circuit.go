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

type circuit struct {
	ID        int             `csv:"circuitId"`
	Ref       string          `csv:"circuitRef"`
	Name      string          `csv:"name"`
	Location  string          `csv:"location"`
	Country   string          `csv:"country"`
	Latitude  decimal.Decimal `csv:"lat"`
	Longitude decimal.Decimal `csv:"lng"`
	Altitude  optionalNumber  `csv:"alt"`
	URL       string          `csv:"url"`
}

func (i Imager) loadCircuits(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/circuits.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening circuits CSV file: %w", err)
	}
	defer file.Close()

	var circuits []*circuit

	if err = gocsv.UnmarshalFile(file, &circuits); err != nil {
		return fmt.Errorf("error marshaling circuits CSV file: %w", err)
	}

	records := []database.SaveCircuitsParams{}

	for _, circuit := range circuits {
		records = append(records, database.SaveCircuitsParams{
			ID:       int32(circuit.ID),
			Ref:      circuit.Ref,
			Name:     circuit.Name,
			Location: &circuit.Location,
			Country:  &circuit.Country,
			Lat:      &circuit.Latitude,
			Lng:      &circuit.Longitude,
			Alt:      circuit.Altitude.Value,
			Url:      circuit.URL,
		})
	}

	_, err = i.db.WithTx(tx).SaveCircuits(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving circuits: %w", err)
	}

	fmt.Println("[Circuits] seeding complete")
	return nil
}
