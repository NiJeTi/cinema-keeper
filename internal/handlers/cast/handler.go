package cast

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/commands"
	"github.com/nijeti/cinema-keeper/internal/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
	"github.com/nijeti/cinema-keeper/internal/types"
)

type Handler struct {
	ctx   context.Context
	log   *slog.Logger
	utils discordUtils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
) Handler {
	log = log.With("command", commands.CastName)

	return Handler{
		ctx:   ctx,
		log:   log,
		utils: discordUtils.New(log),
	}
}

func (h Handler) Handle(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	channelID, err := h.getChannelID(s, i)
	if err != nil && !errors.Is(err, discordgo.ErrStateNotFound) {
		h.log.Error("failed to get channel ID", "error", err)
		return
	}

	channelUsers, err := h.utils.GetVoiceChannelUsers(
		s, types.ID(i.GuildID), types.ID(channelID),
	)
	if err != nil {
		h.log.Error("failed to get channel users", "error", err)
		return
	}

	if (len(channelUsers) == 1 && channelUsers[0].User.ID == i.Member.User.ID) ||
		len(channelUsers) == 0 {
		_ = h.utils.Respond(h.ctx, s, i, responses.CastNoUsers())
	}

	_ = h.utils.Respond(
		h.ctx, s, i, responses.CastUsers(i.Member, channelUsers),
	)
}

func (h Handler) getChannelID(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) (string, error) {
	optionsMap := discordUtils.OptionsMap(i)

	if opt, ok := optionsMap[commands.CastOptionChannel]; ok {
		return opt.ChannelValue(s).ID, nil
	}

	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)

	switch {
	case err == nil:
		return voiceState.ChannelID, nil
	case errors.Is(err, discordgo.ErrStateNotFound):
		_ = h.utils.Respond(h.ctx, s, i, responses.UserNotInVoiceChannel())
		return "", err
	default:
		return "", errors.New("failed to get user voice state")
	}
}
