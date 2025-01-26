package mentionVoiceChan_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/mentionVoiceChan"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
	"github.com/nijeti/cinema-keeper/internal/services/mentionVoiceChan"
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
		channelID *models.DiscordID
	}
	type setup func(
		t *testing.T, args args, err error,
	) *mentionVoiceChan.Service

	tests := map[string]struct {
		args  args
		err   error
		setup setup
	}{
		"voice_state_error": {
			args: args{
				channelID: nil,
			},
			err: errors.New("voice state error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(nil, err)

				return mentionVoiceChan.New(d)
			},
		},
		"no_channel_response_error": {
			args: args{
				channelID: nil,
			},
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.CastNoChannel(),
				).Return(err)

				return mentionVoiceChan.New(d)
			},
		},
		"voice_channel_users_error": {
			args: args{
				channelID: nil,
			},
			err: errors.New("voice channel users error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().VoiceChannelUsers(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(vs.ChannelID),
				).Return(nil, err)

				return mentionVoiceChan.New(d)
			},
		},
		"response_error": {
			args: args{
				channelID: nil,
			},
			err: errors.New("response error"),
			setup: func(
				t *testing.T, _ args, err error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().VoiceChannelUsers(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(vs.ChannelID),
				).Return(nil, nil)

				d.EXPECT().Respond(ctx, i, responses.CastNoUsers()).Return(err)

				return mentionVoiceChan.New(d)
			},
		},
		"no_channel": {
			args: args{
				channelID: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(nil, nil)

				d.EXPECT().Respond(
					ctx, i, responses.CastNoChannel(),
				).Return(nil)

				return mentionVoiceChan.New(d)
			},
		},
		"no_users": {
			args: args{
				channelID: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(vs, nil)

				d.EXPECT().VoiceChannelUsers(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(vs.ChannelID),
				).Return(nil, nil)

				d.EXPECT().Respond(ctx, i, responses.CastNoUsers()).Return(nil)

				return mentionVoiceChan.New(d)
			},
		},
		"self_only": {
			args: args{
				channelID: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(vs, nil)

				users := []*discordgo.Member{i.Member}
				d.EXPECT().VoiceChannelUsers(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(vs.ChannelID),
				).Return(users, nil)

				d.EXPECT().Respond(ctx, i, responses.CastNoUsers()).Return(nil)

				return mentionVoiceChan.New(d)
			},
		},
		"success_channel_id_nil": {
			args: args{
				channelID: nil,
			},
			err: nil,
			setup: func(
				t *testing.T, _ args, _ error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				vs := &discordgo.VoiceState{ChannelID: "3"}
				d.EXPECT().UserVoiceState(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(i.Member.User.ID),
				).Return(vs, nil)

				users := []*discordgo.Member{
					i.Member,
					{User: &discordgo.User{ID: "4"}},
				}
				d.EXPECT().VoiceChannelUsers(
					ctx,
					models.DiscordID(i.GuildID),
					models.DiscordID(vs.ChannelID),
				).Return(users, nil)

				d.EXPECT().Respond(
					ctx, i, responses.CastUsers(i.Member, users),
				).Return(nil)

				return mentionVoiceChan.New(d)
			},
		},
		"success": {
			args: args{
				channelID: ptr.To(models.DiscordID("3")),
			},
			err: nil,
			setup: func(
				t *testing.T, args args, _ error,
			) *mentionVoiceChan.Service {
				d := mocks.NewMockDiscord(t)

				users := []*discordgo.Member{
					i.Member,
					{User: &discordgo.User{ID: "4"}},
				}
				d.EXPECT().VoiceChannelUsers(
					ctx, models.DiscordID(i.GuildID), *args.channelID,
				).Return(users, nil)

				d.EXPECT().Respond(
					ctx, i, responses.CastUsers(i.Member, users),
				).Return(nil)

				return mentionVoiceChan.New(d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.args, tt.err)
				err := s.Exec(ctx, i, tt.args.channelID)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
