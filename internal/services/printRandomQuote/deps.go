package printRandomQuote

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type db interface {
	GetRandomQuoteInGuild(
		ctx context.Context, guildID models.ID,
	) (*models.Quote, error)
}

type discord interface {
	GuildMember(
		ctx context.Context, guildID models.ID, userID models.ID,
	) (*discordgo.Member, error)

	Respond(
		ctx context.Context,
		i *discordgo.Interaction,
		response *discordgo.InteractionResponse,
	) error
}
