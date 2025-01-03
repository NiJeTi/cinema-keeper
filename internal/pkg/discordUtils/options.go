package discordUtils

import (
	"github.com/bwmarrin/discordgo"
)

type interaction interface {
	ApplicationCommandData() discordgo.ApplicationCommandInteractionData
}

func OptionsMap(
	i interaction,
) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options

	optionsMap := make(
		map[string]*discordgo.ApplicationCommandInteractionDataOption,
		len(options),
	)
	for _, option := range options {
		optionsMap[option.Name] = option
	}
	return optionsMap
}
