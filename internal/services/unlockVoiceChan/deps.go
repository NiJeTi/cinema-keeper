package unlockVoiceChan

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

	UserVoiceState(
		ctx context.Context, guildID models.ID, userID models.ID,
	) (*discordgo.VoiceState, error)

	ChannelUnsetUserLimit(
		ctx context.Context, channelID models.ID,
	) error
}
