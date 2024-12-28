package lock

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordutils"
)

type Handler struct {
	ctx     context.Context
	log     *slog.Logger
	session *discordgo.Session
	utils   discordutils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
	session *discordgo.Session,
) *Handler {
	log = log.With("command", discord.LockName)

	return &Handler{
		ctx:     ctx,
		log:     log,
		session: session,
		utils:   discordutils.New(ctx, log, session),
	}
}

func (h *Handler) Handle(i *discordgo.InteractionCreate) {
	guild, err := h.session.State.Guild(i.GuildID)
	if err != nil {
		h.log.ErrorContext(h.ctx, "failed to get guild", "error", err)
		return
	}

	channelID := ""
	channels := make(map[string]int)
	for _, vs := range guild.VoiceStates {
		channels[vs.ChannelID]++

		if vs.UserID == i.Member.User.ID {
			channelID = vs.ChannelID
		}
	}

	if channelID == "" {
		_ = h.utils.Respond(i, responses.UserNotInVoiceChannel())
		return
	}

	limit := channels[channelID]
	if opt, ok := discordutils.OptionsMap(i)[discord.LockOptionLimit]; ok {
		limit = int(opt.IntValue())
	}

	channel, err := h.session.ChannelEdit(
		channelID, &discordgo.ChannelEdit{
			UserLimit: limit,
		},
	)
	if err != nil {
		h.log.Error("failed to set user limit", "error", err)
		return
	}

	_ = h.utils.Respond(i, responses.LockedChannel(channel, limit))
}
