package responses

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func LockSpecifyLimit() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Please specify a user limit",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func LockDone(limit int) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Locked channel for **%d** user(s)", limit),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}
