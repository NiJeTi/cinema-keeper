package commands

import (
	"github.com/bwmarrin/discordgo"
)

const (
	CastName          = "cast"
	CastOptionChannel = "channel"
)

func Cast() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        CastName,
		Description: "Mention all users in a voice channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        CastOptionChannel,
				Description: "Channel to mention",
				Type:        discordgo.ApplicationCommandOptionChannel,
				ChannelTypes: []discordgo.ChannelType{
					discordgo.ChannelTypeGuildVoice,
				},
			},
		},
	}
}
