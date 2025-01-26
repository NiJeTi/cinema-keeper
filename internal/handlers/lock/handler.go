package lock

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
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
	ctx context.Context, i *discordgo.Interaction,
) error {
	var limit *int

	optionsMap := discordUtils.OptionsMap(i)
	if opt, ok := optionsMap[commands.LockOptionLimit]; ok {
		limit = ptr.To(int(opt.IntValue()))
	}

	err := h.service.Exec(ctx, i, limit)
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
