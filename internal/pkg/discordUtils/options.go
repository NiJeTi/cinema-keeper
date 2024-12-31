package discordUtils

import (
	"github.com/bwmarrin/discordgo"
)

func OptionsMap(
	i *discordgo.Interaction,
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
