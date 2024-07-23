package discord

import (
	"github.com/bwmarrin/discordgo"
)

const (
	LockName        = "lock"
	LockOptionLimit = "limit"
)

var (
	lockOptionLimitMinValue float64 = 1
	lockOptionLimitMaxValue float64 = 99
)

var Lock = &discordgo.ApplicationCommand{
	Name:        LockName,
	Description: "Set user limit for current voice channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        LockOptionLimit,
			Description: "Custom limit value",
			Type:        discordgo.ApplicationCommandOptionInteger,
			MinValue:    &lockOptionLimitMinValue,
			MaxValue:    lockOptionLimitMaxValue,
		},
	},
}
