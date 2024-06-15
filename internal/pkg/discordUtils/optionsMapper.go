package discordUtils

import (
	"github.com/bwmarrin/discordgo"
)

func OptionsMap(
	i *discordgo.InteractionCreate,
) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options

	if len(options) == 0 {
		return make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	}

	optionsMap := make(
		map[string]*discordgo.ApplicationCommandInteractionDataOption,
		len(options),
	)
	for _, option := range options {
		optionsMap[option.Name] = option
	}
	return optionsMap
}
