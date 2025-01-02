package lockVoiceChan_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/lockVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
	"github.com/nijeti/cinema-keeper/internal/services/lockVoiceChan"
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

	type args struct {
		limit *int
	}
	type setup func(t *testing.T, args args, err error) *lockVoiceChan.Service

	tests := map[string]struct {
		args  args
		err   error
		setup setup
	}{
		"voice_state_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("voice state error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(nil, err)

				return lockVoiceChan.New(d)
			},
		},
		"voice_state_nil_response_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.UserNotInVoiceChannel(),
				).Return(err)

				return lockVoiceChan.New(d)
			},
		},
		"voice_channel_users_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("voice channel users error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(nil, err)

				return lockVoiceChan.New(d)
			},
		},
		"channel_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("channel error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make([]*discordgo.Member, 1)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(nil, err)

				return lockVoiceChan.New(d)
			},
		},
		"too_many_users_response_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make(
					[]*discordgo.Member, commands.LockOptionLimitMaxValue+1,
				)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().Respond(
					ctx, i, responses.LockSpecifyLimit(c),
				).Return(err)

				return lockVoiceChan.New(d)
			},
		},
		"edit_channel_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("edit channel error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make([]*discordgo.Member, 1)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().EditChannel(
					ctx,
					models.ID(c.ID),
					&discordgo.ChannelEdit{UserLimit: len(users)},
				).Return(err)

				return lockVoiceChan.New(d)
			},
		},
		"response_error": {
			args: args{
				limit: nil,
			},
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make([]*discordgo.Member, 1)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().EditChannel(
					ctx,
					models.ID(c.ID),
					&discordgo.ChannelEdit{UserLimit: len(users)},
				).Return(nil)

				d.EXPECT().Respond(
					ctx, i, responses.LockDone(c, len(users)),
				).Return(err)

				return lockVoiceChan.New(d)
			},
		},
		"voice_state_nil": {
			args: args{
				limit: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.UserNotInVoiceChannel(),
				).Return(nil)

				return lockVoiceChan.New(d)
			},
		},
		"too_many_users": {
			args: args{
				limit: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make(
					[]*discordgo.Member, commands.LockOptionLimitMaxValue+1,
				)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().Respond(
					ctx, i, responses.LockSpecifyLimit(c),
				).Return(nil)

				return lockVoiceChan.New(d)
			},
		},
		"success_limit_nil": {
			args: args{
				limit: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				users := make([]*discordgo.Member, 1)
				d.EXPECT().VoiceChannelUsers(
					ctx, models.ID(i.GuildID), models.ID(vs.ChannelID),
				).Return(users, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().EditChannel(
					ctx,
					models.ID(c.ID),
					&discordgo.ChannelEdit{UserLimit: len(users)},
				).Return(nil)

				d.EXPECT().Respond(
					ctx, i, responses.LockDone(c, len(users)),
				).Return(nil)

				return lockVoiceChan.New(d)
			},
		},
		"success": {
			args: args{
				limit: ptr.To(1),
			},
			err: nil,
			setup: func(
				t *testing.T, args args, _ error,
			) *lockVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx, models.ID(i.GuildID), models.ID(i.Member.User.ID),
				).Return(vs, nil)

				c := &discordgo.Channel{ID: vs.ChannelID}
				d.EXPECT().Channel(
					ctx, models.ID(vs.ChannelID),
				).Return(c, nil)

				d.EXPECT().EditChannel(
					ctx,
					models.ID(c.ID),
					&discordgo.ChannelEdit{UserLimit: *args.limit},
				).Return(nil)

				d.EXPECT().Respond(
					ctx, i, responses.LockDone(c, *args.limit),
				).Return(nil)

				return lockVoiceChan.New(d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.args, tt.err)
				err := s.Exec(ctx, i, tt.args.limit)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
