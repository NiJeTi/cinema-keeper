package unlockVoiceChan_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/unlockVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/unlockVoiceChan"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
		Member: &discordgo.Member{
			User: &discordgo.User{
				ID: "2",
			},
		},
	}

	type setup func(t *testing.T, err error) *unlockVoiceChan.Service

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"voice_state_error": {
			err: errors.New("voice state error"),
			setup: func(
				t *testing.T, err error,
			) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(nil, err)

				return unlockVoiceChan.New(d)
			},
		},
		"voice_state_nil_response_error": {
			err: errors.New("response error"),
			setup: func(t *testing.T, err error) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.UserNotInVoiceChannel(),
				).Return(err)

				return unlockVoiceChan.New(d)
			},
		},
		"channel_limit_unset_error": {
			err: errors.New("channel limit unset error"),
			setup: func(t *testing.T, err error) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().ChannelUnsetUserLimit(
					ctx, models.ID(vs.ChannelID),
				).Return(err)

				return unlockVoiceChan.New(d)
			},
		},
		"respond_error": {
			err: errors.New("channel error"),
			setup: func(t *testing.T, err error) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().ChannelUnsetUserLimit(
					ctx, models.ID(vs.ChannelID),
				).Return(nil)

				d.EXPECT().Respond(ctx, i, responses.UnlockDone()).Return(err)

				return unlockVoiceChan.New(d)
			},
		},
		"voice_state_nil": {
			err: nil,
			setup: func(t *testing.T, _ error) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.UserNotInVoiceChannel(),
				).Return(nil)

				return unlockVoiceChan.New(d)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *unlockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.ID(i.GuildID),
					models.ID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().ChannelUnsetUserLimit(
					ctx, models.ID(vs.ChannelID),
				).Return(nil)

				d.EXPECT().Respond(ctx, i, responses.UnlockDone()).Return(nil)

				return unlockVoiceChan.New(d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.err)
				err := s.Exec(ctx, i)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
