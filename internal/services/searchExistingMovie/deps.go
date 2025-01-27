package searchExistingMovie

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

type db interface {
	GuildMoviesByTitle(
		ctx context.Context, guildID models.DiscordID, title string, limit int,
	) ([]models.MovieBase, error)
}
