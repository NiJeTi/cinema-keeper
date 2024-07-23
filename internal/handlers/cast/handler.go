package cast

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
	"github.com/nijeti/cinema-keeper/internal/types"
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
	log = log.With("command", discord.CastName)

	return &Handler{
		ctx:     ctx,
		log:     log,
		session: session,
		utils:   discordUtils.New(ctx, log, session),
	}
}

func (h *Handler) Handle(i *discordgo.InteractionCreate) {
	channelID, err := h.getChannelID(i)
	if err != nil && !errors.Is(err, discordgo.ErrStateNotFound) {
		h.log.Error("failed to get channel ID", "error", err)
		return
	}

	channelUsers, err := h.utils.GetVoiceChannelUsers(
		types.ID(i.GuildID), types.ID(channelID),
	)
	if err != nil {
		h.log.Error("failed to get channel users", "error", err)
		return
	}

	if (len(channelUsers) == 1 && channelUsers[0].User.ID == i.Member.User.ID) ||
		len(channelUsers) == 0 {
		_ = h.utils.Respond(i, responses.CastNoUsers())
	} else {
		_ = h.utils.Respond(i, responses.CastUsers(i.Member, channelUsers))
	}
}

func (h *Handler) getChannelID(i *discordgo.InteractionCreate) (string, error) {
	optionsMap := discordUtils.OptionsMap(i)

	if opt, ok := optionsMap[discord.CastOptionChannel]; ok {
		return opt.ChannelValue(h.session).ID, nil
	}

	voiceState, err := h.session.State.VoiceState(i.GuildID, i.Member.User.ID)

	switch {
	case err == nil:
		return voiceState.ChannelID, nil
	case errors.Is(err, discordgo.ErrStateNotFound):
		_ = h.utils.Respond(i, responses.UserNotInVoiceChannel())
		return "", err
	default:
		return "", errors.New("failed to get user voice state")
	}
}
