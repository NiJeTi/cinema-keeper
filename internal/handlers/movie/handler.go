package movie

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

type Handler struct {
	searchNewMovie searchNewMovie
	addMovie       addMovie
}

func New(
	searchNewMovie searchNewMovie,
	addMovie addMovie,
) *Handler {
	return &Handler{
		searchNewMovie: searchNewMovie,
		addMovie:       addMovie,
	}
}

func (h *Handler) Handle(
	ctx context.Context, i *discordgo.Interaction,
) error {
	optionsMap := discordUtils.OptionsMap(i)

	if opt, ok := optionsMap[commands.MovieSubCommandAdd]; ok {
		return h.add(ctx, i, opt)
	}

	return nil
}

func (h *Handler) add(
	ctx context.Context,
	i *discordgo.Interaction,
	opt *discordgo.ApplicationCommandInteractionDataOption,
) error {
	optionsMap := discordUtils.SubOptionsMap(opt)
	optValue := optionsMap[commands.MovieOptionTitle].StringValue()

	switch i.Type {
	case discordgo.InteractionApplicationCommandAutocomplete:
		return h.addAutocomplete(ctx, i, optValue)
	case discordgo.InteractionApplicationCommand:
		return h.addExec(ctx, i, optValue)
	default:
		return nil
	}
}

func (h *Handler) addAutocomplete(
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
	if err := h.searchNewMovie.Exec(ctx, i, title); err != nil {
		return fmt.Errorf("failed to autocomplete movie title: %w", err)
	}
	return nil
}

func (h *Handler) addExec(
	ctx context.Context, i *discordgo.Interaction, title string,
) error {
	if err := h.addMovie.Exec(ctx, i, title); err != nil {
		return fmt.Errorf("failed to add movie: %w", err)
	}
	return nil
}
