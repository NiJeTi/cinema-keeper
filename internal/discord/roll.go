package discord

import (
	"sort"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/die"
	"github.com/nijeti/cinema-keeper/internal/pkg/utils"
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

const (
	rollOptionCountMinValue float64 = 1
	rollOptionCountMaxValue float64 = 10
)

//nolint:gochecknoglobals // pending rework
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

	sort.Slice(
		choices, func(i, j int) bool {
			//nolint:errcheck // value is assigned above
			return choices[i].Value.(die.Size) < choices[j].Value.(die.Size)
		},
	)

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
				MinValue:    utils.Ptr(rollOptionCountMinValue),
				MaxValue:    rollOptionCountMaxValue,
			},
		},
	}
}
