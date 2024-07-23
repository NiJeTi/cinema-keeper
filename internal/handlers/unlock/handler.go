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
	ctx     context.Context
	log     *slog.Logger
	session *discordgo.Session
	utils   discordUtils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
	session *discordgo.Session,
) *Handler {
	log = log.With("command", discord.UnlockName)

	return &Handler{
		ctx:     ctx,
		log:     log,
		session: session,
		utils:   discordUtils.New(ctx, log, session),
	}
}

func (h *Handler) Handle(i *discordgo.InteractionCreate) {
	voiceState, err := h.session.State.VoiceState(i.GuildID, i.Member.User.ID)
	if err != nil && !errors.Is(err, discordgo.ErrStateNotFound) {
		h.log.ErrorContext(h.ctx, "failed to get voice state", "error", err)
		return
	}

	if errors.Is(err, discordgo.ErrStateNotFound) {
		_ = h.utils.Respond(i, responses.UserNotInVoiceChannel())
		return
	}

	channel, err := h.unlockChannel(voiceState.ChannelID)
	if err != nil {
		h.log.Error("failed to remove user limit", "error", err)
		return
	}

	_ = h.utils.Respond(i, responses.UnlockedChannel(channel))
}

// unlockChannel Workaround method for removing channel user limit
func (h *Handler) unlockChannel(channelID string) (*discordgo.Channel, error) {
	type request struct {
		UserLimit int `json:"user_limit"`
	}
	resp, err := h.session.RequestWithBucketID(
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
