package cast

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type service interface {
	Exec(
		ctx context.Context,
		i *discordgo.Interaction,
		channelID *models.DiscordID,
	) error
}
