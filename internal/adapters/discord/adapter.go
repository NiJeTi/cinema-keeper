package discord

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type Adapter struct {
	session *discordgo.Session
}

func New(session *discordgo.Session) *Adapter {
	return &Adapter{
		session: session,
	}
}

func (a *Adapter) Respond(
	ctx context.Context,
	i *discordgo.Interaction,
	response *discordgo.InteractionResponse,
) error {
	if i.Type == discordgo.InteractionMessageComponent {
		response.Type = discordgo.InteractionResponseUpdateMessage
	}

	err := a.session.InteractionRespond(
		i, response, discordgo.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("failed to respond to interaction: %w", err)
	}

	return nil
}

func (a *Adapter) GuildMember(
	ctx context.Context, guildID models.DiscordID, userID models.DiscordID,
) (*discordgo.Member, error) {
	member, err := a.session.GuildMember(
		string(guildID), string(userID), discordgo.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild member: %w", err)
	}

	return member, nil
}

func (a *Adapter) Channel(
	ctx context.Context, channelID models.DiscordID,
) (*discordgo.Channel, error) {
	channel, err := a.session.Channel(
		string(channelID), discordgo.WithContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	return channel, nil
}

func (a *Adapter) UserVoiceState(
	_ context.Context, guildID models.DiscordID, userID models.DiscordID,
) (*discordgo.VoiceState, error) {
	voiceState, err := a.session.State.VoiceState(
		string(guildID), string(userID),
	)

	switch {
	case err == nil:
		return voiceState, nil
	case errors.Is(err, discordgo.ErrStateNotFound):
		return nil, nil //nolint:nilnil // nil channel DiscordID is a valid result
	default:
		return nil, fmt.Errorf("failed to get user voice state: %w", err)
	}
}

func (a *Adapter) VoiceChannelUsers(
	ctx context.Context, guildID models.DiscordID, channelID models.DiscordID,
) ([]*discordgo.Member, error) {
	var users []*discordgo.Member

	var after string
	for {
		const maxMembers = 1000
		members, err := a.session.GuildMembers(
			string(guildID), after, maxMembers, discordgo.WithContext(ctx),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get guild members: %w", err)
		}

		if len(members) == 0 {
			break
		}

		for _, member := range members {
			voiceState, err := a.session.State.VoiceState(
				string(guildID), member.User.ID,
			)
			if err != nil {
				if errors.Is(err, discordgo.ErrStateNotFound) {
					continue
				}
				return nil, fmt.Errorf("failed to get voice state: %w", err)
			}

			if voiceState.ChannelID != string(channelID) {
				continue
			}

			users = append(users, member)
		}

		after = members[len(members)-1].User.ID
	}

	return users, nil
}

func (a *Adapter) EditChannel(
	ctx context.Context, channelID models.DiscordID,
	edit *discordgo.ChannelEdit,
) error {
	_, err := a.session.ChannelEdit(
		string(channelID), edit, discordgo.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("failed to edit channel: %w", err)
	}

	return nil
}

// ChannelUnsetUserLimit
// Use this method to remove user limit.
// Package "discordgo" has a flaw which disallows to do it with EditChannel.
func (a *Adapter) ChannelUnsetUserLimit(
	ctx context.Context, channelID models.DiscordID,
) error {
	type request struct {
		UserLimit int `json:"user_limit"`
	}

	_, err := a.session.RequestWithBucketID(
		http.MethodPatch,
		discordgo.EndpointChannel(string(channelID)),
		request{UserLimit: 0},
		discordgo.EndpointChannel(string(channelID)),
		discordgo.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("channel unlock request failed: %w", err)
	}

	return nil
}

func (a *Adapter) SetActivity(
	_ context.Context, activity *discordgo.Activity,
) error {
	err := a.session.UpdateStatusComplex(
		discordgo.UpdateStatusData{
			Status:     string(discordgo.StatusOnline),
			Activities: []*discordgo.Activity{activity},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}
