package addQuote

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
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
	ctx context.Context,
	i *discordgo.Interaction,
	authorID models.DiscordID,
	text string,
) error {
	author, err := s.discord.GuildMember(
		ctx, models.DiscordID(i.GuildID), authorID,
	)
	if err != nil {
		return fmt.Errorf("failed to get author: %w", err)
	}

	quote := &models.Quote{
		AuthorID:  models.DiscordID(author.User.ID),
		Text:      text,
		GuildID:   models.DiscordID(i.GuildID),
		AddedByID: models.DiscordID(i.Member.User.ID),
		Timestamp: time.Now(),
	}
	err = s.db.AddUserQuoteInGuild(ctx, quote)
	if err != nil {
		return fmt.Errorf("failed to save quote: %w", err)
	}

	err = s.discord.Respond(ctx, i, responses.QuoteAdded(author, quote))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
