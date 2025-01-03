package db

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type QuotesRepo struct {
	logger *slog.Logger
	db     *sqlx.DB
}

type quoteRow struct {
	Timestamp time.Time `db:"timestamp"`
	AuthorID  string    `db:"author_id"`
	AddedByID string    `db:"added_by_id"`
	GuildID   string    `db:"guild_id"`
	Text      string    `db:"text"`
}

func NewQuotesRepo(
	logger *slog.Logger,
	db *sqlx.DB,
) *QuotesRepo {
	return &QuotesRepo{
		logger: logger.With("repo", "quotes"),
		db:     db,
	}
}

func (r *QuotesRepo) GetUserQuotesInGuild(
	ctx context.Context, guildID models.ID, authorID models.ID,
) ([]*models.Quote, error) {
	const query = `
		select author_id, text, guild_id, added_by_id, timestamp from quotes
		where guild_id = $1 and author_id = $2
		order by timestamp`

	var rows []quoteRow
	err := r.db.SelectContext(ctx, &rows, query, guildID, authorID)
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
	ctx context.Context, guildID models.ID,
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

	return withTransaction(
		ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
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
		Author: &discordgo.Member{
			User: &discordgo.User{
				ID: row.AuthorID,
			},
		},
		Text:    row.Text,
		GuildID: models.ID(row.GuildID),
		AddedBy: &discordgo.Member{
			User: &discordgo.User{
				ID: row.AddedByID,
			},
		},
		Timestamp: row.Timestamp,
	}
}

func (*QuotesRepo) fromModel(model *models.Quote) quoteRow {
	return quoteRow{
		AuthorID:  model.Author.User.ID,
		Text:      model.Text,
		GuildID:   model.GuildID.String(),
		AddedByID: model.AddedBy.User.ID,
		Timestamp: model.Timestamp,
	}
}
