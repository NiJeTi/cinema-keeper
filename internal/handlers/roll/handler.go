package roll

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

type Handler struct {
	service service
}

func New(
	service service,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(
	ctx context.Context, i *discordgo.InteractionCreate,
) error {
	size := commands.RollOptionSizeDefault
	count := commands.RollOptionCountDefault

	optionsMap := discordUtils.OptionsMap(i.Interaction)
	if opt, ok := optionsMap[commands.RollOptionSize]; ok {
		size = dice.Size(opt.IntValue())
	}
	if opt, ok := optionsMap[commands.RollOptionCount]; ok {
		count = int(opt.IntValue())
	}

	err := h.service.Exec(ctx, i.Interaction, size, count)
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
