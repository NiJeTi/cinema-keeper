package addMovie

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type discord interface {
	Respond(
		ctx context.Context,
		i *discordgo.Interaction,
		response *discordgo.InteractionResponse,
	) error
}

type omdb interface {
	MovieByID(ctx context.Context, id models.ImdbID) (*models.MovieMeta, error)
}

type db interface {
	MovieByImdbID(
		ctx context.Context, id models.ImdbID,
	) (*models.ID, *models.MovieMeta, error)

	AddMovie(
		ctx context.Context, movie *models.MovieMeta,
	) (models.ID, error)

	GuildMovieExists(
		ctx context.Context, id models.ID, guildID models.DiscordID,
	) (bool, error)

	AddMovieToGuild(
		ctx context.Context,
		movieID models.ID,
		guildID models.DiscordID,
		addedByID models.DiscordID,
	) error
}
