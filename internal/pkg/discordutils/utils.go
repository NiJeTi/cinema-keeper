package discordutils

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type Utils interface {
	GetVoiceChannelUsers(
		guild models.ID,
		channel models.ID,
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
	guild models.ID,
	channel models.ID,
) ([]*discordgo.Member, error) {
	guildID := guild.String()
	channelID := channel.String()

	var users []*discordgo.Member

	var after string
	for {
		const maxMembers = 1000
		members, err := u.session.GuildMembers(guildID, after, maxMembers)
		if err != nil {
			return nil, fmt.Errorf("failed to get guild members: %w", err)
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
				return nil, fmt.Errorf("failed to get voice state: %w", err)
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
	return fmt.Errorf("failed to respond: %w", err)
}
