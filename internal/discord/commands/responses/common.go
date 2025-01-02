package responses

import (
	"github.com/bwmarrin/discordgo"
)

func Activity() *discordgo.Activity {
	return &discordgo.Activity{
		Name: "Collecting movies and quotes",
		Type: discordgo.ActivityTypeCustom,
	}
}

func UserNotInVoiceChannel() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You must be in a voice channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
