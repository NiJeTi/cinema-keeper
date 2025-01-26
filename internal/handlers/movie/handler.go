package movie

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (*Handler) Handle(
	_ context.Context, _ *discordgo.Interaction,
) error {
	return nil
}
