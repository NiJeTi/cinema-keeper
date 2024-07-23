package unlock

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
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
	log = log.With("command", discord.UnlockName)

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
	voiceState, err := s.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil && !errors.Is(err, discordgo.ErrStateNotFound) {
		h.log.ErrorContext(h.ctx, "failed to get voice state", "error", err)
		return
	}

	if errors.Is(err, discordgo.ErrStateNotFound) {
		_ = h.utils.Respond(h.ctx, s, i, responses.UserNotInVoiceChannel())
		return
	}

	channel, err := h.unlockChannel(s, voiceState.ChannelID)
	if err != nil {
		h.log.Error("failed to remove user limit", "error", err)
		return
	}

	_ = h.utils.Respond(h.ctx, s, i, responses.UnlockedChannel(channel))
}

func (h Handler) unlockChannel(
	s *discordgo.Session,
	channelID string,
) (*discordgo.Channel, error) {
	type request struct {
		UserLimit int `json:"user_limit"`
	}
	resp, err := s.RequestWithBucketID(
		http.MethodPatch,
		discordgo.EndpointChannel(channelID),
		request{UserLimit: 0},
		discordgo.EndpointChannel(channelID),
	)
	if err != nil {
		return nil, err
	}

	channel := &discordgo.Channel{}
	err = discordgo.Unmarshal(resp, channel)
	return channel, err
}
