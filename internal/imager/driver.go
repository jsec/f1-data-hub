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

type driver struct {
	ID          int32          `csv:"driverId"`
	Ref         string         `csv:"driverRef"`
	Number      optionalNumber `csv:"number"`
	Code        optionalString `csv:"code"`
	FirstName   string         `csv:"forename"`
	LastName    string         `csv:"surname"`
	DOB         dateOnly       `csv:"dob"`
	Nationality string         `csv:"nationality"`
	URL         string         `csv:"url"`
}

func (i imager) loadDrivers(ctx context.Context, tx pgx.Tx) error {
	spinner := i.spinners.AddSpinner("Seeding drivers")

	file, err := os.OpenFile("data/drivers.csv", os.O_RDONLY, 0600)
	if err != nil {
		spinner.Error()
		return fmt.Errorf("error opening driver CSV file: %w", err)
	}
	defer file.Close()

	var drivers []*driver

	if err = gocsv.UnmarshalFile(file, &drivers); err != nil {
		spinner.Error()
		return fmt.Errorf("error marshaling driver CSV file: %w", err)
	}

	records := []database.SaveDriversParams{}

	for _, driver := range drivers {
		records = append(records, database.SaveDriversParams{
			ID:          driver.ID,
			Ref:         driver.Ref,
			Number:      driver.Number.Value,
			Code:        driver.Code.Value,
			FirstName:   driver.FirstName,
			LastName:    driver.LastName,
			DateOfBirth: &driver.DOB.Value,
			Nationality: &driver.Nationality,
			Url:         driver.URL,
		})
	}

	if err = i.driverService.SeedDrivers(ctx, tx, records); err != nil {
		spinner.Error()
		return fmt.Errorf("error saving drivers: %w", err)
	}

	spinner.Complete()
	return nil
}

type driverStanding struct {
	ID           int32           `csv:"driverStandingsId"`
	RaceID       int32           `csv:"raceId"`
	DriverID     int32           `csv:"driverId"`
	Points       decimal.Decimal `csv:"points"`
	Position     int32           `csv:"position"`
	PositionText string          `csv:"positionText"`
	Wins         int32           `csv:"wins"`
}

func (i imager) loadDriverStandings(ctx context.Context, tx pgx.Tx) error {
	spinner := i.spinners.AddSpinner("Seeding driver standings")

	file, err := os.OpenFile("data/driver_standings.csv", os.O_RDONLY, 0600)
	if err != nil {
		spinner.Error()
		return fmt.Errorf("error opening driver standings CSV file: %w", err)
	}
	defer file.Close()

	var standings []*driverStanding

	if err = gocsv.UnmarshalFile(file, &standings); err != nil {
		spinner.Error()
		return fmt.Errorf("error marshaling driver standings CSV file: %w", err)
	}

	records := []database.SaveDriverStandingsParams{}

	for _, standing := range standings {
		records = append(records, database.SaveDriverStandingsParams{
			ID:       standing.ID,
			RaceID:   standing.RaceID,
			DriverID: standing.DriverID,
			Points:   standing.Points,
			Position: &standing.Position,
			PosText:  &standing.PositionText,
			Wins:     &standing.Wins,
		})
	}

	if err = i.driverService.SeedDriverStandings(ctx, tx, records); err != nil {
		spinner.Error()
		return fmt.Errorf("error saving driver standings: %w", err)
	}

	spinner.Complete()
	return nil
}
