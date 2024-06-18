package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/die"
)

const (
	RollName        = "roll"
	RollOptionSize  = "size"
	RollOptionCount = "count"
)

const (
	RollOptionSizeDefault  = die.D100
	RollOptionCountDefault = 1
)

var (
	rollOptionCountMinValue float64 = 1
	rollOptionCountMaxValue float64 = 10
)

var Roll = buildRoll()

func buildRoll() *discordgo.ApplicationCommand {
	dieSizes := die.Sizes()
	choices := make(
		[]*discordgo.ApplicationCommandOptionChoice, 0, len(dieSizes),
	)
	for t, s := range dieSizes {
		c := &discordgo.ApplicationCommandOptionChoice{
			Name:  t.String(),
			Value: s,
		}
		choices = append(choices, c)
	}

	return &discordgo.ApplicationCommand{
		Name:        RollName,
		Description: "Roll a die",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        RollOptionSize,
				Description: "Size of a die",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices:     choices,
			},
			{
				Name:        RollOptionCount,
				Description: "Number of dice",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &rollOptionCountMinValue,
				MaxValue:    rollOptionCountMaxValue,
			},
		},
	}
}
