package movie

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type searchNewMovies interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}

type searchGuildMovies interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}

type addMovie interface {
	Exec(ctx context.Context, i *discordgo.Interaction, title string) error
}
