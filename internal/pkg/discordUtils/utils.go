package discordUtils

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type Utils interface {
	GetVoiceChannelUsers(
		guild types.ID,
		channel types.ID,
	) ([]*discordgo.Member, error)

	Respond(
		interaction *discordgo.InteractionCreate,
		response *discordgo.InteractionResponse,
	) error
}

type utils struct {
	ctx     context.Context
	log     *slog.Logger
	session *discordgo.Session
}

func New(
	ctx context.Context,
	log *slog.Logger,
	session *discordgo.Session,
) Utils {
	return &utils{
		ctx:     ctx,
		log:     log,
		session: session,
	}
}

func (u *utils) GetVoiceChannelUsers(
	guild types.ID,
	channel types.ID,
) ([]*discordgo.Member, error) {
	guildID := guild.String()
	channelID := channel.String()

	var users []*discordgo.Member

	var after string
	for {
		members, err := u.session.GuildMembers(guildID, after, 1000)
		if err != nil {
			return nil, err
		}

		if len(members) == 0 {
			break
		}

		for _, member := range members {
			voiceState, err := u.session.State.VoiceState(
				guildID, member.User.ID,
			)
			if err != nil {
				if errors.Is(err, discordgo.ErrStateNotFound) {
					continue
				}
				return nil, err
			}

			if voiceState.ChannelID != channelID {
				continue
			}

			users = append(users, member)
		}

		after = members[len(members)-1].User.ID
	}

	return users, nil
}

func (u *utils) Respond(
	interaction *discordgo.InteractionCreate,
	response *discordgo.InteractionResponse,
) error {
	err := u.session.InteractionRespond(interaction.Interaction, response)
	if err != nil {
		u.log.ErrorContext(u.ctx, "failed to respond", "error", err)
	}
	return err
}
