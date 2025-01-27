package searchExistingMovie

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
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
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
	if len(title) < commands.MovieTitleAutocompleteMinLength {
		return s.respondEmptySearch(ctx, i)
	}

	movies, err := s.db.GuildMoviesByTitle(
		ctx,
		models.DiscordID(i.GuildID),
		title,
		commands.MovieTitleAutocompleteChoicesLimit,
	)
	if err != nil {
		return fmt.Errorf("failed to search movies: %w", err)
	}

	if len(movies) == 0 {
		return s.respondEmptySearch(ctx, i)
	}

	err = s.discord.Respond(ctx, i, responses.MovieAutocompleteSearch(movies))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) respondEmptySearch(
	ctx context.Context, i *discordgo.Interaction,
) error {
	err := s.discord.Respond(
		ctx, i, responses.MovieAutocompleteEmptySearch(),
	)
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
