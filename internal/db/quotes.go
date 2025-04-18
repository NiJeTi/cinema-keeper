package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type QuotesRepo struct {
	repoBase
}

type quoteRow struct {
	ID        int64     `db:"id"`
	AuthorID  string    `db:"author_id"`
	Text      string    `db:"text"`
	GuildID   string    `db:"guild_id"`
	AddedByID string    `db:"added_by_id"`
	Timestamp time.Time `db:"timestamp"`
}

func NewQuotesRepo(
	db *sqlx.DB,
) *QuotesRepo {
	return &QuotesRepo{
		repoBase: repoBase{
			db: db,
		},
	}
}

func (r *QuotesRepo) CountUserQuotesInGuild(
	ctx context.Context, guildID models.DiscordID, authorID models.DiscordID,
) (int, error) {
	const query = `
		select count(*) from quotes
		where guild_id = $1 and author_id = $2`

	var rows []int
	err := r.db.SelectContext(ctx, &rows, query, guildID, authorID)
	if err != nil {
		return 0, fmt.Errorf(
			"failed to select user quotes count in guild: %w", err,
		)
	}

	return rows[0], nil
}

func (r *QuotesRepo) GetUserQuotesInGuild(
	ctx context.Context,
	guildID models.DiscordID,
	authorID models.DiscordID,
	offset, limit int,
) ([]*models.Quote, error) {
	if offset < 0 {
		panic("offset must be greater than or equal to 0")
	}
	if limit <= 0 {
		panic("limit must be greater than 0")
	}

	const query = `
		select author_id, text, guild_id, added_by_id, timestamp from quotes
		where guild_id = $1 and author_id = $2
		order by timestamp desc
		offset $3 limit $4`

	var rows []quoteRow
	err := r.db.SelectContext(
		ctx, &rows, query, guildID, authorID, offset, limit,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to select user quotes in guild: %w", err,
		)
	}

	quotes := make([]*models.Quote, 0, len(rows))
	for _, row := range rows {
		quotes = append(quotes, r.toModel(row))
	}
	return quotes, nil
}

func (r *QuotesRepo) GetRandomQuoteInGuild(
	ctx context.Context, guildID models.DiscordID,
) (*models.Quote, error) {
	const query = `
		select author_id, text, guild_id, added_by_id, timestamp from quotes
		where guild_id = $1
		order by random()
		limit 1`

	var rows []quoteRow
	err := r.db.SelectContext(ctx, &rows, query, guildID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to select random quote in guild: %w", err,
		)
	}

	if len(rows) == 0 {
		return nil, nil //nolint:nilnil // nil is valid if 0 rows were selected
	}

	return r.toModel(rows[0]), nil
}

func (r *QuotesRepo) AddUserQuoteInGuild(
	ctx context.Context, quote *models.Quote,
) error {
	const query = `
		insert into quotes(author_id, text, guild_id, added_by_id, timestamp)
		values ($1, $2, $3, $4, $5)`

	row := r.fromModel(quote)

	return r.withTransaction(
		ctx, func(ctx context.Context, tx *sqlx.Tx) error {
			_, err := tx.ExecContext(
				ctx,
				query,
				row.AuthorID,
				row.Text,
				row.GuildID,
				row.AddedByID,
				row.Timestamp,
			)
			if err != nil {
				return fmt.Errorf("failed to insert quote: %w", err)
			}

			return nil
		},
	)
}

func (*QuotesRepo) toModel(row quoteRow) *models.Quote {
	return &models.Quote{
		AuthorID:  models.DiscordID(row.AuthorID),
		Text:      row.Text,
		GuildID:   models.DiscordID(row.GuildID),
		AddedByID: models.DiscordID(row.AddedByID),
		Timestamp: row.Timestamp,
	}
}

func (*QuotesRepo) fromModel(model *models.Quote) quoteRow {
	return quoteRow{
		AuthorID:  string(model.AuthorID),
		Text:      model.Text,
		GuildID:   string(model.GuildID),
		AddedByID: string(model.AddedByID),
		Timestamp: model.Timestamp,
	}
}
