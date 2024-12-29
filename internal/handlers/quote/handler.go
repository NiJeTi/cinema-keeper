package quote

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordutils"
)

type Handler struct {
	ctx     context.Context
	log     *slog.Logger
	session *discordgo.Session
	db      db
	utils   discordutils.Utils
}

func New(
	ctx context.Context,
	log *slog.Logger,
	session *discordgo.Session,
	db db,
) *Handler {
	log = log.With("command", discord.QuoteName)

	return &Handler{
		ctx:     ctx,
		log:     log,
		session: session,
		db:      db,
		utils:   discordutils.New(ctx, log, session),
	}
}

func (h *Handler) Handle(i *discordgo.InteractionCreate) {
	author, text, err := h.getOptions(i)
	if err != nil {
		h.log.ErrorContext(h.ctx, "failed to get options", "error", err)
	}

	if text == "" {
		h.listQuotes(i, *author)
	} else {
		h.addQuote(i, *author, text)
	}
}

func (h *Handler) getOptions(
	i *discordgo.InteractionCreate,
) (*discordgo.Member, string, error) {
	optionsMap := discordutils.OptionsMap(i)

	authorID := optionsMap[discord.QuoteOptionAuthor].UserValue(h.session).ID
	author, err := h.session.State.Member(i.GuildID, authorID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get author: %w", err)
	}

	var text string
	if opt, ok := optionsMap[discord.QuoteOptionText]; ok {
		text = opt.StringValue()
	}

	return author, text, nil
}
