package responses

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func UnlockedChannel(channel *discordgo.Channel) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Unclocked %s", channel.Mention()),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
