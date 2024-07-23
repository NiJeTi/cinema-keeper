package quote

import (
	"context"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

type Handler struct {
	ctx   context.Context
	log   *slog.Logger
	utils discordUtils.Utils
	db    db
}

func New(
	ctx context.Context,
	log *slog.Logger,
	db db,
) Handler {
	log = log.With("command", discord.QuoteName)

	return Handler{
		ctx:   ctx,
		log:   log,
		utils: discordUtils.New(log),
		db:    db,
	}
}

func (h Handler) Handle(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	author, text, err := h.getOptions(s, i)
	if err != nil {
		h.log.ErrorContext(h.ctx, "failed to get options", "error", err)
	}

	if text == "" {
		h.listQuotes(s, i, author)
	} else {
		h.addQuote(s, i, author, text)
	}
}

func (h Handler) getOptions(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) (*discordgo.Member, string, error) {
	optionsMap := discordUtils.OptionsMap(i)

	authorID := optionsMap[discord.QuoteOptionAuthor].UserValue(s).ID
	author, err := s.State.Member(i.GuildID, authorID)
	if err != nil {
		return nil, "", err
	}

	var text string
	if opt, ok := optionsMap[discord.QuoteOptionText]; ok {
		text = opt.StringValue()
	}

	return author, text, nil
}
