package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/nijeti/cinema-keeper/internal/db/migrations"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbutils"
)

type Migrator struct {
	log       *slog.Logger
	db        *sql.DB
	txWrapper dbutils.TxWrapper
}

func NewMigrator(
	log *slog.Logger,
	db *sql.DB,
	txWrapper dbutils.TxWrapper,
) *Migrator {
	return &Migrator{
		log:       log.With("type", "migrator"),
		db:        db,
		txWrapper: txWrapper,
	}
}

func (m *Migrator) Migrate() error {
	err := m.createMigrationTable()
	if err != nil {
		m.log.Error(
			"failed to create migration table",
			"error", err,
		)
		return err
	}

	for _, migration := range migrations.Migrations() {
		applied, err := m.checkMigration(migration.Name)
		if err != nil {
			m.log.Error(
				"failed to check migration",
				"migration", migration,
				"error", err,
			)
			return err
		}

		if applied {
			m.log.Info(
				"skipping already applied migration",
				"migration", migration,
			)
			continue
		}

		err = m.txWrapper.Wrap(
			context.Background(),
			func(_ context.Context, tx *sql.Tx) error {
				return migration.Execute(tx)
			},
		)
		if err != nil {
			m.log.Error(
				"failed to execute migration",
				"migration", migration.Name,
				"error", err,
			)
			return err
		}

		m.log.Info(
			"applied migration",
			"migration", migration.Name,
		)

		err = m.registerAppliedMigration(migration.Name)
		if err != nil {
			m.log.Error(
				"failed to register applied migration",
				"error", err,
			)
			return err
		}
	}

	return nil
}

func (m *Migrator) createMigrationTable() error {
	return m.txWrapper.Wrap(
		context.Background(), func(_ context.Context, tx *sql.Tx) error {
			_, err := tx.Exec(
				`create table if not exists "migrations"
				(
					"id" serial primary key,
					"name" varchar(32) not null
				)`,
			)
			return err //nolint:wrapcheck // error passthrough
		},
	)
}

func (m *Migrator) registerAppliedMigration(migration string) error {
	return m.txWrapper.Wrap(
		context.Background(), func(_ context.Context, tx *sql.Tx) error {
			_, err := tx.Exec(
				`insert into "migrations"(name) values ($1)`,
				migration,
			)
			return err //nolint:wrapcheck // error passthrough
		},
	)
}

func (m *Migrator) checkMigration(migration string) (bool, error) {
	const query = `select exists(select 1 from "migrations" where name = $1)`

	rows, err := m.db.Query(query, migration)
	if err != nil {
		return false, fmt.Errorf("failed to query db: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return false, errors.New("failed to retrieve rows")
	}
	if rows.Err() != nil {
		return false, fmt.Errorf("failed to read row: %w", rows.Err())
	}

	var applied bool
	err = rows.Scan(&applied)
	if err != nil {
		return false, fmt.Errorf("failed to scan row: %w", err)
	}

	return applied, nil
}
