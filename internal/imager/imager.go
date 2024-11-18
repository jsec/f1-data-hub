package imager

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
)

type Imager struct {
	db *database.Queries
}

func Run(ctx context.Context) error {
	pool, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Error acquiring transaction: %w", err)
	}

	imager := Imager{
		db: database.New(pool),
	}

	return imager.Seed(ctx, tx)
}

func (i Imager) Seed(ctx context.Context, tx pgx.Tx) error {
	loaders := []func(ctx context.Context, tx pgx.Tx) error{
		i.loadSeasons,
		i.loadStatuses,
		i.loadDrivers,
		i.loadConstructors,
		i.loadCircuits,
		i.loadRaces,
		i.loadDriverStandings,
		i.loadConstructorStandings,
		i.loadLapTimes,
		i.loadPitStops,
		i.loadQualifying,
		i.loadResults,
		i.loadConstructorResults,
		i.loadSprintResults,
	}

	for _, loader := range loaders {
		err := loader(ctx, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
