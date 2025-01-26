package discordUtils

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var ErrUnknownInteractionType = errors.New("unknown interaction type")

func ParseCommand(
	i *discordgo.Interaction,
) (command string, subCommand string, err error) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand,
		discordgo.InteractionApplicationCommandAutocomplete:
		return parseInteraction(i.ApplicationCommandData())
	case discordgo.InteractionMessageComponent:
		return parseComponentInteraction(i.MessageComponentData())
	default:
		return "", "", ErrUnknownInteractionType
	}
}

func MustParseCommand(
	i *discordgo.Interaction,
) (command string, subCommand string) {
	cmd, subCmd, err := ParseCommand(i)
	if err != nil {
		panic(err)
	}

	return cmd, subCmd
}

func parseInteraction(
	i discordgo.ApplicationCommandInteractionData,
) (command string, subCommand string, err error) {
	command = i.Name

	for _, opt := range i.Options {
		if opt.Type == discordgo.ApplicationCommandOptionSubCommand {
			subCommand = opt.Name
			break
		}
	}

	return command, subCommand, nil
}

func parseComponentInteraction(
	i discordgo.MessageComponentInteractionData,
) (command string, subCommand string, err error) {
	return ParseCustomID(i.CustomID)
}
