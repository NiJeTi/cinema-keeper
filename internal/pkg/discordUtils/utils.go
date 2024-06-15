package discordUtils

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type Utils struct {
	log *slog.Logger
}

func New(
	log *slog.Logger,
) Utils {
	return Utils{
		log: log,
	}
}

func (u Utils) GetVoiceChannelUsers(
	s *discordgo.Session,
	guild types.ID,
	channel types.ID,
) ([]*discordgo.Member, error) {
	guildID := guild.String()
	channelID := channel.String()

	var users []*discordgo.Member

	var after string
	for {
		members, err := s.GuildMembers(guildID, after, 1000)
		if err != nil {
			return nil, err
		}

		if len(members) == 0 {
			break
		}

		for _, member := range members {
			voiceState, err := s.State.VoiceState(guildID, member.User.ID)
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

func (u Utils) Respond(
	ctx context.Context,
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	response *discordgo.InteractionResponse,
) error {
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		u.log.ErrorContext(ctx, "failed to respond", "error", err)
	}
	return err
}
