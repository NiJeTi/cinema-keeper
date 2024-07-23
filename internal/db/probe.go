package db

import (
	"context"
	"database/sql"
)

type Probe struct {
	db *sql.DB
}

func NewProbe(db *sql.DB) Probe {
	return Probe{
		db: db,
	}
}

func (p Probe) Check(ctx context.Context) error {
	return p.db.PingContext(ctx)
}
