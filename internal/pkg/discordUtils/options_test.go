package discordUtils_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

func TestOptionsMap(t *testing.T) {
	t.Parallel()

	wantOptions := map[string]*discordgo.ApplicationCommandInteractionDataOption{
		"string": {
			Name: "string",
			Type: discordgo.ApplicationCommandOptionString,
		},
		"int": {
			Name: "int",
			Type: discordgo.ApplicationCommandOptionInteger,
		},
		"bool": {
			Name: "bool",
			Type: discordgo.ApplicationCommandOptionBoolean,
		},
	}

	data := discordgo.ApplicationCommandInteractionData{}
	for _, opt := range wantOptions {
		data.Options = append(data.Options, opt)
	}

	options := discordUtils.OptionsMap(
		&discordgo.Interaction{
			Type: discordgo.InteractionApplicationCommand,
			Data: data,
		},
	)

	assert.Equal(t, wantOptions, options)
}

func TestSubOptionsMap(t *testing.T) {
	t.Parallel()

	wantOptions := map[string]*discordgo.ApplicationCommandInteractionDataOption{
		"string": {
			Name: "string",
			Type: discordgo.ApplicationCommandOptionString,
		},
		"int": {
			Name: "int",
			Type: discordgo.ApplicationCommandOptionInteger,
		},
		"bool": {
			Name: "bool",
			Type: discordgo.ApplicationCommandOptionBoolean,
		},
	}

	optArg := &discordgo.ApplicationCommandInteractionDataOption{}
	for _, opt := range wantOptions {
		optArg.Options = append(optArg.Options, opt)
	}

	options := discordUtils.SubOptionsMap(optArg)

	assert.Equal(t, wantOptions, options)
}
