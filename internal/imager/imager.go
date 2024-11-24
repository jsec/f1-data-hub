package imager

import (
	"context"

	"github.com/chelnak/ysmrr"
	"github.com/jackc/pgx/v5"
	"github.com/jsec/f1-data-hub/internal/database"
	"github.com/jsec/f1-data-hub/internal/services"
)

type imager struct {
	spinners           ysmrr.SpinnerManager
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

func newImager(db *database.Queries) imager {
	return imager{
		spinners:           ysmrr.NewSpinnerManager(),
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
}

func (i imager) seed(ctx context.Context, tx pgx.Tx) error {
	i.spinners.Start()

	if err := i.downloadData(); err != nil {
		i.spinners.Stop()
		return err
	}

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
			i.spinners.Stop()
			return err
		}
	}

	i.spinners.Stop()
	return nil
}
