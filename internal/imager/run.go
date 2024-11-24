package imager

import (
	"context"
	"fmt"

	"github.com/jsec/f1-data-hub/internal/database"
)

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

	imager := newImager(database.New(pool))

	if err = imager.seed(ctx, tx); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
