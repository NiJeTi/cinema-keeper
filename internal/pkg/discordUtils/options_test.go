package discordUtils_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/pkg/discordUtils"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

func TestOptionsMap(t *testing.T) {
	t.Parallel()

	i := mocks.NewMockInteraction(t)

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
	i.EXPECT().ApplicationCommandData().Return(data)

	options := discordUtils.OptionsMap(i)

	assert.Equal(t, wantOptions, options)
}
