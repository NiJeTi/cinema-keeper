package discordUtils_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

func TestParseCommand(t *testing.T) {
	t.Parallel()

	type want struct {
		command    string
		subCommand string
		err        error
	}

	tests := map[string]struct {
		i    *discordgo.Interaction
		want want
	}{
		"unknown_type": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionPing,
			},
			want: want{
				err: discordUtils.ErrUnknownInteractionType,
			},
		},
		"interaction_command_only": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Name: "command",
				},
			},
			want: want{
				command: "command",
			},
		},
		"interaction_command_and_sub": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Name: "command",
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name: "sub",
							Type: discordgo.ApplicationCommandOptionSubCommand,
						},
					},
				},
			},
			want: want{
				command:    "command",
				subCommand: "sub",
			},
		},
		"component_command": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionMessageComponent,
				Data: discordgo.MessageComponentInteractionData{
					CustomID: "command_args",
				},
			},
			want: want{
				command: "command",
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				command, subCommand, err := discordUtils.ParseCommand(tt.i)
				assert.Equal(t, tt.want.command, command)
				assert.Equal(t, tt.want.subCommand, subCommand)
				assert.ErrorIs(t, err, tt.want.err)
			},
		)
	}
}
