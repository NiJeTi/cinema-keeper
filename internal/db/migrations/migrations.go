package migrations

import (
	"database/sql"
)

type Migration struct {
	Name    string
	Execute func(tx *sql.Tx) error
}

func Migrations() []Migration {
	return []Migration{
		{
			Name: "initial",
			Execute: func(tx *sql.Tx) error {
				_, err := tx.Exec(
					`create table if not exists "quotes"
					(
					    "id"          serial primary key,
						"author_id"   varchar(20) not null,
						"text"        text        not null,
						"guild_id"    varchar(20) not null,
						"added_by_id" varchar(20) not null,
						"timestamp"   timestamptz not null
					)`,
				)
				if err != nil {
					return err
				}

				_, err = tx.Exec(
					`create index if not exists "quotes_author_id"
    					on "quotes" using hash ("author_id")`,
				)
				return err
			},
		},
	}
}
