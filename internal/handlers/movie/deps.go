package movie

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type searchNewMovie interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}

type searchExistingMovie interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}

type addMovie interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}
