package responses

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func CastNoChannel() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Either specify a channel or join one",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func CastNoUsers() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "There are no users in the voice channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func CastUsers(
	initiator *discordgo.Member,
	users []*discordgo.Member,
) *discordgo.InteractionResponse {
	mentions := make([]string, 0, len(users))
	for _, user := range users {
		if user.User.ID == initiator.User.ID {
			continue
		}

		mentions = append(mentions, user.User.Mention())
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strings.Join(mentions, " "),
		},
	}
}
