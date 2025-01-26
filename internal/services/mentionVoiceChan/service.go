package mentionVoiceChan

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

type Service struct {
	discord discord
}

func New(
	discord discord,
) *Service {
	return &Service{
		discord: discord,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, channelID *models.DiscordID,
) error {
	channelID, err := s.verifyChannelID(ctx, i, channelID)
	if err != nil {
		return err
	}

	if channelID == nil {
		err = s.discord.Respond(ctx, i, responses.CastNoChannel())
		if err != nil {
			return fmt.Errorf("failed to respond: %w", err)
		}

		return nil
	}

	users, err := s.discord.VoiceChannelUsers(
		ctx, models.DiscordID(i.GuildID), *channelID,
	)
	if err != nil {
		return fmt.Errorf("failed to get voice channel users: %w", err)
	}

	var response *discordgo.InteractionResponse
	if (len(users) == 1 && users[0].User.ID == i.Member.User.ID) ||
		len(users) == 0 {
		response = responses.CastNoUsers()
	} else {
		response = responses.CastUsers(i.Member, users)
	}

	err = s.discord.Respond(ctx, i, response)
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}

func (s *Service) verifyChannelID(
	ctx context.Context, i *discordgo.Interaction, channelID *models.DiscordID,
) (*models.DiscordID, error) {
	if channelID != nil {
		return channelID, nil
	}

	voiceState, err := s.discord.UserVoiceState(
		ctx, models.DiscordID(i.GuildID), models.DiscordID(i.Member.User.ID),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user voice channel: %w", err)
	}

	if voiceState == nil {
		return nil, nil //nolint:nilnil // nil is a valid value
	}
	return ptr.To(models.DiscordID(voiceState.ChannelID)), nil
}
