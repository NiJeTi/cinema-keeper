package quote

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

type Handler struct {
	printRandomQuote printRandomQuote
	listUserQuotes   listUserQuotes
	addQuote         addQuote
}

func New(
	printRandomQuote printRandomQuote,
	listUserQuotes listUserQuotes,
	addQuote addQuote,
) *Handler {
	return &Handler{
		printRandomQuote: printRandomQuote,
		listUserQuotes:   listUserQuotes,
		addQuote:         addQuote,
	}
}

func (h *Handler) Handle(
	ctx context.Context, i *discordgo.InteractionCreate,
) error {
	subCommandsMap := discordUtils.OptionsMap(i.Interaction)

	if sc, ok := subCommandsMap[commands.QuoteSubCommandGet]; ok {
		return h.getSubCommand(ctx, i, sc)
	}
	if sc, ok := subCommandsMap[commands.QuoteSubCommandAdd]; ok {
		return h.addSubCommand(ctx, i, sc)
	}

	return nil
}

func (h *Handler) getSubCommand(
	ctx context.Context,
	i *discordgo.InteractionCreate,
	subCommand *discordgo.ApplicationCommandInteractionDataOption,
) error {
	optionsMap := discordUtils.SubOptionsMap(subCommand)

	var authorID *models.ID
	if opt, ok := optionsMap[commands.QuoteOptionAuthor]; ok {
		authorID = ptr.To(models.ID(opt.UserValue(nil).ID))
	}

	var err error
	if authorID != nil {
		err = h.listUserQuotes.Exec(ctx, i.Interaction, *authorID)
	} else {
		err = h.printRandomQuote.Exec(ctx, i.Interaction)
	}
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}

func (h *Handler) addSubCommand(
	ctx context.Context,
	i *discordgo.InteractionCreate,
	subCommand *discordgo.ApplicationCommandInteractionDataOption,
) error {
	optionsMap := discordUtils.SubOptionsMap(subCommand)

	authorID := models.ID(
		optionsMap[commands.QuoteOptionAuthor].UserValue(nil).ID,
	)
	text := optionsMap[commands.QuoteOptionText].StringValue()

	if err := h.addQuote.Exec(ctx, i.Interaction, authorID, text); err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
