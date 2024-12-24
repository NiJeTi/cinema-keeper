package discord

import (
	"github.com/bwmarrin/discordgo"
)

const (
	CastName          = "cast"
	CastOptionChannel = "channel"
)

//nolint:gochecknoglobals // pending rework
var Cast = &discordgo.ApplicationCommand{
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
