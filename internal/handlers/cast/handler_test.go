package cast_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/handlers/cast"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

func TestHandler_Handle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type setup func(
		t *testing.T, i *discordgo.InteractionCreate, err error,
	) *cast.Handler

	tests := map[string]struct {
		i     *discordgo.InteractionCreate
		err   error
		setup setup
	}{
		"service_error": {
			i: &discordgo.InteractionCreate{
				Interaction: &discordgo.Interaction{
					Type: discordgo.InteractionApplicationCommand,
					Data: discordgo.ApplicationCommandInteractionData{},
				},
			},
			err: errors.New("service error"),
			setup: func(
				t *testing.T, i *discordgo.InteractionCreate, err error,
			) *cast.Handler {
				s := mocks.NewMockService(t)

				var channelID *models.ID
				s.EXPECT().Exec(ctx, i.Interaction, channelID).Return(err)

				return cast.New(s)
			},
		},
		"success_channel_id_nil": {
			i: &discordgo.InteractionCreate{
				Interaction: &discordgo.Interaction{
					Type: discordgo.InteractionApplicationCommand,
					Data: discordgo.ApplicationCommandInteractionData{},
				},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.InteractionCreate, _ error,
			) *cast.Handler {
				s := mocks.NewMockService(t)

				var channelID *models.ID
				s.EXPECT().Exec(ctx, i.Interaction, channelID).Return(nil)

				return cast.New(s)
			},
		},
		"success": {
			i: &discordgo.InteractionCreate{
				Interaction: &discordgo.Interaction{
					Type: discordgo.InteractionApplicationCommand,
					Data: discordgo.ApplicationCommandInteractionData{
						Options: []*discordgo.ApplicationCommandInteractionDataOption{
							{
								Name:  commands.CastOptionChannel,
								Type:  discordgo.ApplicationCommandOptionChannel,
								Value: "1",
							},
						},
					},
				},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.InteractionCreate, _ error,
			) *cast.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(
					ctx, i.Interaction, ptr.To(models.ID("1")),
				).Return(nil)

				return cast.New(s)
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
