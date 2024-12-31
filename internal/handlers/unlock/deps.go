package unlock

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type service interface {
	Exec(ctx context.Context, i *discordgo.Interaction) error
}
