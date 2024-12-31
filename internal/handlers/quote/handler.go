package quote

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

type Handler struct {
	listQuotes listQuotes
	addQuote   addQuote
}

func New(
	listQuotes listQuotes,
	addQuote addQuote,
) *Handler {
	return &Handler{
		listQuotes: listQuotes,
		addQuote:   addQuote,
	}
}

func (h *Handler) Handle(
	ctx context.Context, i *discordgo.InteractionCreate,
) error {
	optionsMap := discordUtils.OptionsMap(i.Interaction)

	opt := optionsMap[commands.QuoteOptionAuthor]
	authorID := models.ID(opt.UserValue(nil).ID)

	var text string
	if opt, ok := optionsMap[commands.QuoteOptionText]; ok {
		text = opt.StringValue()
	}

	var err error
	if text == "" {
		err = h.listQuotes.Exec(ctx, i.Interaction, authorID)
	} else {
		err = h.addQuote.Exec(ctx, i.Interaction, authorID, text)
	}
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
