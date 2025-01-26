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
		ctx context.Context,
		guildID models.DiscordID,
		userID models.DiscordID,
	) (*discordgo.VoiceState, error)

	VoiceChannelUsers(
		ctx context.Context,
		guildID models.DiscordID,
		channelID models.DiscordID,
	) ([]*discordgo.Member, error)
}
