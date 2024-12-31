package listQuotes

import (
	"context"
	"fmt"

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
	ctx context.Context, i *discordgo.Interaction, authorID models.ID,
) error {
	author, err := s.discord.GuildMember(ctx, models.ID(i.GuildID), authorID)
	if err != nil {
		return fmt.Errorf("failed to get author: %w", err)
	}

	quotes, err := s.db.GetUserQuotesOnGuild(
		ctx, models.ID(i.GuildID), authorID,
	)
	if err != nil {
		return fmt.Errorf("failed to get quotes: %w", err)
	}

	if len(quotes) == 0 {
		return s.respondEmpty(ctx, i, author)
	}

	return s.respondList(ctx, i, author, quotes)
}

func (s *Service) respondEmpty(
	ctx context.Context, i *discordgo.Interaction, author *discordgo.Member,
) error {
	err := s.discord.Respond(ctx, i, responses.QuotesEmpty(author))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) respondList(
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

	err := s.discord.Respond(ctx, i, responses.QuotesHeader(author))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	const pageSize = 10
	for offset := 0; offset < len(quotes); offset += pageSize {
		limit := offset + pageSize
		if offset+pageSize >= len(quotes) {
			limit = len(quotes)
		}

		err = s.discord.SendEmbeds(
			ctx, models.ID(i.ChannelID), responses.Quotes(quotes[offset:limit]),
		)
		if err != nil {
			return fmt.Errorf("failed to send quotes page: %w", err)
		}
	}

	return nil
}
