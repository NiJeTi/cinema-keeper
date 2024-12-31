package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

const (
	RollName        = "roll"
	RollOptionSize  = "size"
	RollOptionCount = "count"
)

const (
	RollOptionSizeDefault  = dice.D100
	RollOptionCountDefault = 1
)

const (
	rollOptionCountMinValue float64 = 1
	rollOptionCountMaxValue float64 = 10
)

func Roll() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        RollName,
		Description: "Roll a die",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        RollOptionSize,
				Description: "Size of a die",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{Name: dice.D2Name, Value: dice.D2},
					{Name: dice.D4Name, Value: dice.D4},
					{Name: dice.D6Name, Value: dice.D6},
					{Name: dice.D8Name, Value: dice.D8},
					{Name: dice.D10Name, Value: dice.D10},
					{Name: dice.D12Name, Value: dice.D12},
					{Name: dice.D20Name, Value: dice.D20},
					{Name: dice.D100Name, Value: dice.D100},
				},
			},
			{
				Name:        RollOptionCount,
				Description: "Number of dice",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    ptr.To(rollOptionCountMinValue),
				MaxValue:    rollOptionCountMaxValue,
			},
		},
	}
}
