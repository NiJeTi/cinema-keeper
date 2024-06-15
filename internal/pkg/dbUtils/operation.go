package dbUtils

import (
	"context"
	"database/sql"
)

type Operation func(ctx context.Context, tx *sql.Tx) error
