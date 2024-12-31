package quote

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type listQuotes interface {
	Exec(
		ctx context.Context, i *discordgo.Interaction, authorID models.ID,
	) error
}

type addQuote interface {
	Exec(
		ctx context.Context,
		i *discordgo.Interaction,
		authorID models.ID,
		text string,
	) error
}
