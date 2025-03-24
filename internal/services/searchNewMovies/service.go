package searchNewMovies

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
)

type Service struct {
	discord discord
	omdb    omdb
}

func New(
	discord discord,
	omdb omdb,
) *Service {
	return &Service{
		discord: discord,
		omdb:    omdb,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
	if len(title) < commands.MovieTitleAutocompleteMinLength {
		return s.respondEmptySearch(ctx, i)
	}

	movies, err := s.omdb.MoviesByTitle(ctx, title)
	if err != nil {
		return fmt.Errorf("failed to search movies: %w", err)
	}

	if len(movies) == 0 {
		return s.respondEmptySearch(ctx, i)
	}

	if len(movies) > commands.MovieTitleAutocompleteChoicesLimit {
		movies = movies[:commands.MovieTitleAutocompleteChoicesLimit]
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
