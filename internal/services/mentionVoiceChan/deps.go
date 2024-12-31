package mentionVoiceChan

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

	VoiceChannelUsers(
		ctx context.Context, guildID models.ID, channelID models.ID,
	) ([]*discordgo.Member, error)
}
