package searchNewMovies

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
	MoviesByTitle(
		ctx context.Context, title string,
	) ([]models.MovieBase, error)
}
