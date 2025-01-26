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
	ctx context.Context, i *discordgo.Interaction,
) error {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		return h.handlerNew(ctx, i)
	case discordgo.InteractionMessageComponent:
		return h.handleContinue(ctx, i)
	default:
		return nil
	}
}

func (h *Handler) handlerNew(
	ctx context.Context, i *discordgo.Interaction,
) error {
	optionsMap := discordUtils.OptionsMap(i)

	if opt, ok := optionsMap[commands.QuoteSubCommandGet]; ok {
		return h.getSubCommand(ctx, i, opt)
	}
	if opt, ok := optionsMap[commands.QuoteSubCommandAdd]; ok {
		return h.addSubCommand(ctx, i, opt)
	}

	return nil
}

func (h *Handler) handleContinue(
	ctx context.Context, i *discordgo.Interaction,
) error {
	_, subCommand := discordUtils.MustParseCommand(i)

	//nolint:revive,gocritic // extension-ready approach
	switch subCommand {
	case commands.QuoteSubCommandGet:
		return h.getSubCommandContinue(ctx, i)
	}

	return nil
}

func (h *Handler) getSubCommand(
	ctx context.Context,
	i *discordgo.Interaction,
	opt *discordgo.ApplicationCommandInteractionDataOption,
) error {
	subOptionsMap := discordUtils.SubOptionsMap(opt)

	var authorID *models.ID
	if opt, ok := subOptionsMap[commands.QuoteOptionAuthor]; ok {
		authorID = ptr.To(models.ID(opt.UserValue(nil).ID))
	}

	if authorID == nil {
		if err := h.printRandomQuote.Exec(ctx, i); err != nil {
			return fmt.Errorf("failed to print random quote: %w", err)
		}
		return nil
	}

	if err := h.listUserQuotes.Exec(ctx, i, *authorID, 0); err != nil {
		return fmt.Errorf("failed to list user quotes: %w", err)
	}
	return nil
}

func (h *Handler) addSubCommand(
	ctx context.Context,
	i *discordgo.Interaction,
	subCommand *discordgo.ApplicationCommandInteractionDataOption,
) error {
	optionsMap := discordUtils.SubOptionsMap(subCommand)

	authorID := models.ID(optionsMap[commands.QuoteOptionAuthor].UserValue(nil).ID)
	text := optionsMap[commands.QuoteOptionText].StringValue()

	if err := h.addQuote.Exec(ctx, i, authorID, text); err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}

func (h *Handler) getSubCommandContinue(
	ctx context.Context, i *discordgo.Interaction,
) error {
	authorID, page := commands.QuoteParseButtonCustomID(
		i.MessageComponentData().CustomID,
	)

	if err := h.listUserQuotes.Exec(ctx, i, authorID, page); err != nil {
		return fmt.Errorf("failed to update user quotes list: %w", err)
	}

	return nil
}
