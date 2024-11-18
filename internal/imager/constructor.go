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

type constructor struct {
	ID          int    `csv:"constructorId"`
	Ref         string `csv:"constructorRef"`
	Name        string `csv:"name"`
	Nationality string `csv:"nationality"`
	URL         string `csv:"url"`
}

func (i Imager) loadConstructors(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/constructors.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening constructors CSV file: %w", err)
	}
	defer file.Close()

	var constructors []*constructor

	if err = gocsv.UnmarshalFile(file, &constructors); err != nil {
		return fmt.Errorf("error marshaling constructors CSV file: %w", err)
	}

	records := []database.SaveConstructorsParams{}

	for _, constructor := range constructors {
		records = append(records, database.SaveConstructorsParams{
			ID:          int32(constructor.ID),
			Ref:         constructor.Ref,
			Name:        constructor.Name,
			Nationality: &constructor.Nationality,
			Url:         constructor.URL,
		})
	}

	_, err = i.db.WithTx(tx).SaveConstructors(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving constructors: %w", err)
	}

	return nil
}

type constructorResult struct {
	ID            int32           `csv:"constructorResultsId"`
	RaceID        int32           `csv:"raceId"`
	ConstructorID int32           `csv:"constructorId"`
	Points        decimal.Decimal `csv:"points"`
	Status        optionalString  `csv:"status"`
}

func (i Imager) loadConstructorResults(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/constructor_results.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening constructor results CSV file: %w", err)
	}
	defer file.Close()

	var results []*constructorResult

	if err = gocsv.UnmarshalFile(file, &results); err != nil {
		return fmt.Errorf("error marshaling constructor results CSV file: %w", err)
	}

	records := []database.SaveConstructorResultsParams{}

	for _, result := range results {
		records = append(records, database.SaveConstructorResultsParams{
			ID:            result.ID,
			RaceID:        result.RaceID,
			ConstructorID: result.ConstructorID,
			Points:        &result.Points,
			Status:        result.Status.Value,
		})
	}

	_, err = i.db.WithTx(tx).SaveConstructorResults(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving constructor results: %w", err)
	}

	fmt.Println("[Constructor Results] seeding complete")
	return nil
}

type constructorStanding struct {
	ID            int32           `csv:"constructorStandingsId"`
	RaceID        int32           `csv:"raceId"`
	ConstructorID int32           `csv:"constructorId"`
	Points        decimal.Decimal `csv:"points"`
	Position      int32           `csv:"position"`
	PositionText  string          `csv:"positionText"`
	Wins          int32           `csv:"wins"`
}

func (i Imager) loadConstructorStandings(ctx context.Context, tx pgx.Tx) error {
	file, err := os.OpenFile("data/constructor_standings.csv", os.O_RDONLY, 0600)
	if err != nil {
		return fmt.Errorf("error opening constructor standings CSV file: %w", err)
	}
	defer file.Close()

	var standings []*constructorStanding

	if err = gocsv.UnmarshalFile(file, &standings); err != nil {
		return fmt.Errorf("error marshaling constructor standings CSV file: %w", err)
	}

	records := []database.SaveConstructorStandingsParams{}

	for _, standing := range standings {
		records = append(records, database.SaveConstructorStandingsParams{
			ID:            standing.ID,
			RaceID:        standing.RaceID,
			ConstructorID: standing.ConstructorID,
			Points:        standing.Points,
			Position:      &standing.Position,
			PosText:       &standing.PositionText,
			Wins:          standing.Wins,
		})
	}

	_, err = i.db.WithTx(tx).SaveConstructorStandings(ctx, records)
	if err != nil {
		return fmt.Errorf("error saving constructor standings: %w", err)
	}

	fmt.Println("[Constructor Standings] seeding complete")
	return nil
}
