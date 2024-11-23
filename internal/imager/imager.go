package imager

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
	"github.com/jsec/f1-data-hub/internal/services"
)

type Imager struct {
	circuitService     services.CircuitService
	constructorService services.ConstructorService
	driverService      services.DriverService
	lapTimeService     services.LapTimeService
	pitStopService     services.PitStopService
	qualifyingService  services.QualifyingService
	raceService        services.RaceService
	resultService      services.ResultsService
	seasonService      services.SeasonService
	statusService      services.StatusService
}

func Run(ctx context.Context) error {
	pool, err := database.Connect(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring database connection: %w", err)
	}
	defer pool.Close()

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	db := database.New(pool)

	imager := Imager{
		circuitService:     services.NewCircuitService(db),
		constructorService: services.NewConstructorService(db),
		driverService:      services.NewDriverService(db),
		lapTimeService:     services.NewLapTimeService(db),
		pitStopService:     services.NewPitStopService(db),
		qualifyingService:  services.NewQualifyingService(db),
		raceService:        services.NewRaceService(db),
		resultService:      services.NewResultsService(db),
		seasonService:      services.NewSeasonService(db),
		statusService:      services.NewStatusService(db),
	}

	if err = imager.Seed(ctx, tx); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
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
