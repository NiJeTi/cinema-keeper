package searchGuildMovies

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
	if title == "" {
		return s.returnLastMovies(ctx, i)
	}
	return s.returnMoviesByTitle(ctx, i, title)
}

func (s *Service) returnLastMovies(
	ctx context.Context, i *discordgo.Interaction,
) error {
	movies, err := s.db.GuildMovies(
		ctx,
		models.DiscordID(i.GuildID),
		0,
		commands.MovieTitleAutocompleteChoicesLimit,
	)
	if err != nil {
		return fmt.Errorf("failed to get guild movies: %w", err)
	}

	if len(movies) == 0 {
		return s.respondEmpty(ctx, i)
	}
	return s.respondMovies(ctx, i, movies)
}

func (s *Service) returnMoviesByTitle(
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
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
		return s.respondEmpty(ctx, i)
	}
	return s.respondMovies(ctx, i, movies)
}

func (s *Service) respondEmpty(
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

func (s *Service) respondMovies(
	ctx context.Context, i *discordgo.Interaction, movies []models.MovieBase,
) error {
	err := s.discord.Respond(ctx, i, responses.MovieAutocompleteSearch(movies))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
