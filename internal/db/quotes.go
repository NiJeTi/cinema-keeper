package db

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/dbUtils"
	"github.com/nijeti/cinema-keeper/internal/types"
)

type QuotesRepo struct {
	log       *slog.Logger
	db        *sql.DB
	txWrapper dbUtils.TxWrapper
}

func NewQuotesRepo(
	log *slog.Logger,
	db *sql.DB,
	txWrapper dbUtils.TxWrapper,
) QuotesRepo {
	return QuotesRepo{
		log:       log.With("repo", "quotes"),
		db:        db,
		txWrapper: txWrapper,
	}
}

func (r QuotesRepo) GetUserQuotesOnGuild(
	ctx context.Context,
	authorID types.ID,
	guildID types.ID,
) ([]*models.Quote, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`select author_id, text, guild_id, added_by_id, timestamp from quotes
			where author_id = $1 and guild_id = $2
			order by timestamp`,
		authorID.String(), guildID.String(),
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to get user quotes", "error", err)
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote
	for rows.Next() {
		quote := &models.Quote{}
		err = rows.Scan(
			&quote.AuthorID,
			&quote.Text,
			&quote.GuildID,
			&quote.AddedByID,
			&quote.Timestamp,
		)
		if err != nil {
			r.log.ErrorContext(ctx, "failed to parse quote", "error", err)
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	err = rows.Err()
	if err != nil {
		r.log.ErrorContext(ctx, "failed to fetch quotes", "error", err)
		return nil, err
	}

	return quotes, nil
}

func (r QuotesRepo) AddUserQuoteOnGuild(
	ctx context.Context,
	quote *models.Quote,
) error {
	quote.Timestamp = time.Now().UTC()

	err := r.txWrapper.Wrap(
		ctx, func(ctx context.Context, tx *sql.Tx) error {
			_, err := tx.ExecContext(
				ctx,
				`insert into quotes(author_id, text, guild_id, added_by_id, timestamp)
				values ($1, $2, $3, $4, $5)`,
				quote.AuthorID.String(),
				quote.Text,
				quote.GuildID.String(),
				quote.AddedByID.String(),
				quote.Timestamp,
			)

			return err
		},
	)
	if err != nil {
		r.log.ErrorContext(ctx, "failed to insert quote", "error", err)
	}
	return err
}
