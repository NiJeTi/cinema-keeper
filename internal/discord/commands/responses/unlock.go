package responses

import (
	"github.com/bwmarrin/discordgo"
)

func UnlockDone() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Unlocked channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
