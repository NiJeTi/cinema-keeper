package lock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/handlers/lock"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

func TestHandler_Handle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type setup func(
		t *testing.T, i *discordgo.Interaction, err error,
	) *lock.Handler

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
			) *lock.Handler {
				s := mocks.NewMockService(t)

				var limit *int
				s.EXPECT().Exec(ctx, i, limit).Return(err)

				return lock.New(s)
			},
		},
		"success_limit_nil": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.Interaction, _ error,
			) *lock.Handler {
				s := mocks.NewMockService(t)

				var limit *int
				s.EXPECT().Exec(ctx, i, limit).Return(nil)

				return lock.New(s)
			},
		},
		"success": {
			i: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name:  commands.LockOptionLimit,
							Type:  discordgo.ApplicationCommandOptionInteger,
							Value: float64(1),
						},
					},
				},
			},
			err: nil,
			setup: func(
				t *testing.T, i *discordgo.Interaction, _ error,
			) *lock.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(ctx, i, ptr.To(1)).Return(nil)

				return lock.New(s)
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
