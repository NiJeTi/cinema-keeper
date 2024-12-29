package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func withTransaction(
	ctx context.Context,
	db *sqlx.DB,
	op func(ctx context.Context, tx *sqlx.Tx) error,
) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = op(ctx, tx)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf(
				"failed to execute transaction: %w; rollback error: %w",
				err, txErr,
			)
		}

		return fmt.Errorf("failed to execute transaction: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
