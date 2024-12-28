package dbutils

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type TxWrapper interface {
	Wrap(ctx context.Context, operation Operation) error
}

type txWrapper struct {
	log *slog.Logger
	db  *sql.DB
}

func NewTxWrapper(
	log *slog.Logger,
	db *sql.DB,
) TxWrapper {
	return &txWrapper{
		log: log.With("type", "transaction wrapper"),
		db:  db,
	}
}

func (w *txWrapper) Wrap(
	ctx context.Context,
	operation Operation,
) error {
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		w.log.ErrorContext(ctx, "failed to begin transaction", "error", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = operation(ctx, tx)
	if err != nil {
		w.rollbackTx(ctx, tx)
		return fmt.Errorf("failed to execute operation: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		w.log.ErrorContext(ctx, "failed to commit transaction", "error", err)

		w.rollbackTx(ctx, tx)

		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (w *txWrapper) rollbackTx(ctx context.Context, tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		w.log.ErrorContext(ctx, "failed to rollback transaction", "error", err)
	}
}
