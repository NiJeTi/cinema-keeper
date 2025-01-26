package unlock

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
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
	err := h.service.Exec(ctx, i)
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
