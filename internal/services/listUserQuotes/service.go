package listUserQuotes

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/paginator"
)

type Service struct {
	discord discord
	db      db
}

func New(
	discord discord,
	db db,
) *Service {
	return &Service{
		discord: discord,
		db:      db,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, authorID models.ID, page int,
) error {
	author, err := s.discord.GuildMember(ctx, models.ID(i.GuildID), authorID)
	if err != nil {
		return fmt.Errorf("failed to get author: %w", err)
	}

	quoteCount, err := s.db.CountUserQuotesInGuild(
		ctx, models.ID(i.GuildID), authorID,
	)
	if err != nil {
		return fmt.Errorf("failed to count quotes: %w", err)
	}

	if quoteCount == 0 {
		return s.respondEmpty(ctx, i, author)
	}

	return s.respondList(ctx, i, author, quoteCount, page)
}

func (s *Service) respondEmpty(
	ctx context.Context, i *discordgo.Interaction, author *discordgo.Member,
) error {
	err := s.discord.Respond(ctx, i, responses.QuoteUserNoQuotes(author))
	if err != nil {
		return fmt.Errorf("failed to send empty response: %w", err)
	}

	return nil
}

func (s *Service) respondList(
	ctx context.Context,
	i *discordgo.Interaction,
	author *discordgo.Member,
	quoteCount int,
	page int,
) error {
	offset := page * commands.QuoteMaxQuotesPerPage

	quotes, err := s.db.GetUserQuotesInGuild(
		ctx,
		models.ID(i.GuildID),
		models.ID(author.User.ID),
		offset,
		commands.QuoteMaxQuotesPerPage,
	)
	if err != nil {
		return fmt.Errorf("failed to get quotes: %w", err)
	}

	err = s.enrichQuotes(ctx, i, author, quotes)
	if err != nil {
		return err
	}

	pageInfo := paginator.Info(
		quoteCount, commands.QuoteMaxQuotesPerPage, offset,
	)
	err = s.discord.Respond(
		ctx, i, responses.QuoteList(quotes, pageInfo.Page, pageInfo.LastPage),
	)
	if err != nil {
		return fmt.Errorf("failed to send quotes page: %w", err)
	}

	return nil
}

func (s *Service) enrichQuotes(
	ctx context.Context,
	i *discordgo.Interaction,
	author *discordgo.Member,
	quotes []*models.Quote,
) error {
	for _, q := range quotes {
		addedBy, err := s.discord.GuildMember(
			ctx, models.ID(i.GuildID), models.ID(q.AddedBy.User.ID),
		)
		if err != nil {
			return fmt.Errorf("failed to get guild member: %w", err)
		}

		q.Author = author
		q.AddedBy = addedBy
	}
	return nil
}
