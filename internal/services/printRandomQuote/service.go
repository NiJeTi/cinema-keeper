package printRandomQuote

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
)

type Service struct {
	db      db
	discord discord
}

func New(
	db db,
	discord discord,
) *Service {
	return &Service{
		db:      db,
		discord: discord,
	}
}

func (s *Service) Exec(ctx context.Context, i *discordgo.Interaction) error {
	quote, err := s.db.GetRandomQuoteInGuild(ctx, models.DiscordID(i.GuildID))
	if err != nil {
		return fmt.Errorf("failed to get random quote: %w", err)
	}

	if quote == nil {
		return s.respondNoQuotes(ctx, i)
	}

	return s.respondQuote(ctx, i, quote)
}

func (s *Service) respondNoQuotes(
	ctx context.Context, i *discordgo.Interaction,
) error {
	err := s.discord.Respond(ctx, i, responses.QuoteGuildNoQuotes())
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) respondQuote(
	ctx context.Context, i *discordgo.Interaction, quote *models.Quote,
) error {
	author, err := s.discord.GuildMember(
		ctx, models.DiscordID(i.GuildID), quote.AuthorID,
	)
	if err != nil {
		return fmt.Errorf("failed to get author: %w", err)
	}

	err = s.discord.Respond(ctx, i, responses.QuoteRandomQuote(author, quote))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
