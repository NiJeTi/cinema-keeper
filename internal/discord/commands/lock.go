package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

const (
	LockName        = "lock"
	LockOptionLimit = "limit"
)

const LockOptionLimitMaxValue = 99

const (
	lockOptionLimitMinValue float64 = 1
	lockOptionLimitMaxValue float64 = LockOptionLimitMaxValue
)

func Lock() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        LockName,
		Description: "Set user limit for current voice channel",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        LockOptionLimit,
				Description: "Custom limit value",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    ptr.To(lockOptionLimitMinValue),
				MaxValue:    lockOptionLimitMaxValue,
			},
		},
	}
}
