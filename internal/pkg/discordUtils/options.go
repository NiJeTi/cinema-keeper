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

func SubOptionsMap(
	opt *discordgo.ApplicationCommandInteractionDataOption,
) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	if opt == nil {
		return nil
	}

	optionsMap := make(
		map[string]*discordgo.ApplicationCommandInteractionDataOption,
		len(opt.Options),
	)
	for _, option := range opt.Options {
		optionsMap[option.Name] = option
	}
	return optionsMap
}
