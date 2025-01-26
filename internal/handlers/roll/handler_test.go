package roll_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/handlers/roll"
	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
)

func TestHandler_Handle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type setup func(
		t *testing.T, i *discordgo.Interaction, err error,
	) *roll.Handler

	tests := map[string]struct {
		i     *discordgo.Interaction
		err   error
		setup setup
	}{
		"service_error": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{},
			},
			err: errors.New("service error"),
			setup: func(
				t *testing.T, i *discordgo.Interaction, err error,
			) *roll.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(
					ctx,
					i,
					commands.RollOptionSizeDefault,
					commands.RollOptionCountDefault,
				).Return(err)

				return roll.New(s)
			},
		},
		"success_no_options": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.Interaction, _ error,
			) *roll.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(
					ctx,
					i,
					commands.RollOptionSizeDefault,
					commands.RollOptionCountDefault,
				).Return(nil)

				return roll.New(s)
			},
		},
		"success": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name:  commands.RollOptionSize,
							Type:  discordgo.ApplicationCommandOptionInteger,
							Value: float64(dice.D20),
						},
						{
							Name:  commands.RollOptionCount,
							Type:  discordgo.ApplicationCommandOptionInteger,
							Value: float64(1),
						},
					},
				},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.Interaction, _ error,
			) *roll.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(
					ctx, i, dice.D20, 1,
				).Return(nil)

				return roll.New(s)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				h := tt.setup(t, tt.i, tt.err)
				err := h.Handle(ctx, tt.i)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
