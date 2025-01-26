package addMovie

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
	omdb    omdb
	db      db
}

func New(
	discord discord,
	omdb omdb,
	db db,
) *Service {
	return &Service{
		discord: discord,
		omdb:    omdb,
		db:      db,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
	if ok := commands.MovieValidateImdbID(title); !ok {
		err := s.discord.Respond(ctx, i, responses.MovieInvalidTitle())
		if err != nil {
			return fmt.Errorf("failed to respond: %w", err)
		}
		return nil
	}

	imdbID := models.ImdbID(title)

	movieID, movie, err := s.db.MovieByImdbID(ctx, imdbID)
	if err != nil {
		return fmt.Errorf("failed to get movie from db: %w", err)
	}

	guildID := models.DiscordID(i.GuildID)

	if movieID != nil {
		exists, err := s.checkMovieExists(ctx, i, movieID, guildID)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}
	} else {
		movieID, movie, err = s.addNewMovie(ctx, imdbID)
		if err != nil {
			return err
		}
	}

	addedByID := models.DiscordID(i.Member.User.ID)

	err = s.db.AddMovieToGuild(ctx, *movieID, guildID, addedByID)
	if err != nil {
		return fmt.Errorf("failed to add movie to guild: %w", err)
	}

	err = s.discord.Respond(ctx, i, responses.MovieAdded(movie))
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) checkMovieExists(
	ctx context.Context, i *discordgo.Interaction, movieID *models.ID,
	guildID models.DiscordID,
) (bool, error) {
	exists, err := s.db.GuildMovieExists(ctx, *movieID, guildID)
	if err != nil {
		return false, fmt.Errorf(
			"failed to check if movie already in list: %w", err,
		)
	}

	if !exists {
		return false, nil
	}

	err = s.discord.Respond(ctx, i, responses.MovieAlreadyExists())
	if err != nil {
		return false, fmt.Errorf("failed to respond: %w", err)
	}

	return true, nil
}

func (s *Service) addNewMovie(
	ctx context.Context, id models.ImdbID,
) (*models.ID, *models.MovieMeta, error) {
	movieMeta, err := s.omdb.MovieByID(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get movie from omdb: %w", err)
	}

	movieID, err := s.db.AddMovie(ctx, movieMeta)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to add movie to db: %w", err)
	}

	return &movieID, movieMeta, nil
}
