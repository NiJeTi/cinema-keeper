package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/utils"
)

const (
	LockName        = "lock"
	LockOptionLimit = "limit"
)

const (
	lockOptionLimitMinValue float64 = 1
	lockOptionLimitMaxValue float64 = 99
)

//nolint:gochecknoglobals // pending rework
var Lock = &discordgo.ApplicationCommand{
	Name:        LockName,
	Description: "Set user limit for current voice channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        LockOptionLimit,
			Description: "Custom limit value",
			Type:        discordgo.ApplicationCommandOptionInteger,
			MinValue:    utils.Ptr(lockOptionLimitMinValue),
			MaxValue:    lockOptionLimitMaxValue,
		},
	},
}
