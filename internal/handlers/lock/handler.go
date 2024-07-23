package lock

import (
	"context"
	"log/slog"

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
	log = log.With("command", discord.LockName)

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
	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		h.log.ErrorContext(h.ctx, "failed to get guild", "error", err)
		return
	}

	channelID := ""
	channels := map[string]int{}
	for _, vs := range guild.VoiceStates {
		channels[vs.ChannelID]++

		if vs.UserID == i.Member.User.ID {
			channelID = vs.ChannelID
		}
	}

	if channelID == "" {
		_ = h.utils.Respond(h.ctx, s, i, responses.UserNotInVoiceChannel())
		return
	}

	limit := channels[channelID]
	if opt, ok := discordUtils.OptionsMap(i)[discord.LockOptionLimit]; ok {
		limit = int(opt.IntValue())
	}

	channel, err := s.ChannelEdit(
		channelID, &discordgo.ChannelEdit{
			UserLimit: limit,
		},
	)
	if err != nil {
		h.log.Error("failed to set user limit", "error", err)
		return
	}

	_ = h.utils.Respond(h.ctx, s, i, responses.LockedChannel(channel, limit))
}
