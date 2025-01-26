package cast

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
	var channelID *models.DiscordID

	optionsMap := discordUtils.OptionsMap(i)
	if opt, ok := optionsMap[commands.CastOptionChannel]; ok {
		channelID = ptr.To(models.DiscordID(opt.ChannelValue(nil).ID))
	}

	err := h.service.Exec(ctx, i, channelID)
	if err != nil {
		return fmt.Errorf("failed to execute service: %w", err)
	}

	return nil
}
