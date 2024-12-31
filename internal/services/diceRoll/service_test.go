//nolint:dupl // table-driven tests
package diceRoll_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/diceRoll"
	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
	"github.com/nijeti/cinema-keeper/internal/services/diceRoll"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{}

	type args struct {
		size  dice.Size
		count int
	}
	type setup func(t *testing.T, args args, err error) *diceRoll.Service

	tests := map[string]struct {
		args  args
		err   error
		setup setup
	}{
		"discord_error": {
			args: args{
				size:  commands.RollOptionSizeDefault,
				count: commands.RollOptionCountDefault,
			},
			err: errors.New("discord error"),
			setup: func(t *testing.T, _ args, err error) *diceRoll.Service {
				d := mocks.NewMockDiscord(t)
				d.EXPECT().
					Respond(ctx, mock.Anything, mock.Anything).
					Return(err)

				return diceRoll.New(d)
			},
		},
		"single_default_roll": {
			args: args{
				size:  commands.RollOptionSizeDefault,
				count: commands.RollOptionCountDefault,
			},
			err: nil,
			setup: func(t *testing.T, _ args, err error) *diceRoll.Service {
				d := mocks.NewMockDiscord(t)
				d.EXPECT().Respond(ctx, mock.Anything, mock.Anything).
					RunAndReturn(
						func(
							_ context.Context,
							_ *discordgo.Interaction,
							r *discordgo.InteractionResponse,
						) error {
							assert.NotNil(t, r)
							assert.NotNil(t, r.Data)
							assert.Regexp(t, `^\d{1,3}$`, r.Data.Content)

							return err
						},
					)

				return diceRoll.New(d)
			},
		},
		"single_d20_roll": {
			args: args{
				size:  dice.D20,
				count: commands.RollOptionCountDefault,
			},
			err: nil,
			setup: func(t *testing.T, _ args, err error) *diceRoll.Service {
				d := mocks.NewMockDiscord(t)
				d.EXPECT().Respond(ctx, mock.Anything, mock.Anything).
					RunAndReturn(
						func(
							_ context.Context,
							_ *discordgo.Interaction,
							r *discordgo.InteractionResponse,
						) error {
							assert.NotNil(t, r)
							assert.NotNil(t, r.Data)
							assert.Regexp(
								t, `^\*\*D20\*\*:\s\d{1,2}$`, r.Data.Content,
							)

							return err
						},
					)

				return diceRoll.New(d)
			},
		},
		"multiple_d100_roll": {
			args: args{
				size:  dice.D100,
				count: 10, //nolint:revive // random value
			},
			err: nil,
			setup: func(t *testing.T, args args, err error) *diceRoll.Service {
				d := mocks.NewMockDiscord(t)
				d.EXPECT().Respond(ctx, mock.Anything, mock.Anything).
					RunAndReturn(
						func(
							_ context.Context,
							_ *discordgo.Interaction,
							r *discordgo.InteractionResponse,
						) error {
							assert.NotNil(t, r)
							assert.NotNil(t, r.Data)
							assert.Len(t, r.Data.Embeds, 1)
							assert.Equal(
								t, dice.D100Name, r.Data.Embeds[0].Title,
							)
							assert.Equal(
								t,
								args.count,
								strings.Count(
									r.Data.Embeds[0].Description, ">",
								),
							)

							return err
						},
					)

				return diceRoll.New(d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.args, tt.err)
				err := s.Exec(ctx, i, tt.args.size, tt.args.count)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
