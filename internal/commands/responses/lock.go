package responses

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func LockedChannel(
	channel *discordgo.Channel,
	limit int,
) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"Locked channel %s for **%d** user(s)",
				channel.Mention(), limit,
			),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	}
}
