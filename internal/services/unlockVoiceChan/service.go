package unlockVoiceChan

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
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

func (s *Service) Exec(ctx context.Context, i *discordgo.Interaction) error {
	voiceState, err := s.discord.UserVoiceState(
		ctx, models.DiscordID(i.GuildID), models.DiscordID(i.Member.User.ID),
	)
	if err != nil {
		return fmt.Errorf("failed to get user voice channel: %w", err)
	}

	if voiceState == nil {
		err = s.discord.Respond(ctx, i, responses.UserNotInVoiceChannel())
		if err != nil {
			return fmt.Errorf("failed to respond: %w", err)
		}

		return nil
	}

	err = s.discord.ChannelUnsetUserLimit(
		ctx, models.DiscordID(voiceState.ChannelID),
	)
	if err != nil {
		return fmt.Errorf("failed to unset user limit: %w", err)
	}

	err = s.discord.Respond(ctx, i, responses.UnlockDone())
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
