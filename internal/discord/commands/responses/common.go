package responses

import (
	"github.com/bwmarrin/discordgo"
)

func UserNotInVoiceChannel() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You must be in a voice channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
