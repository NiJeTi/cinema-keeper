package listUserQuotes

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type discord interface {
	Respond(
		ctx context.Context,
		i *discordgo.Interaction,
		response *discordgo.InteractionResponse,
	) error

	GuildMember(
		ctx context.Context, guildID models.ID, userID models.ID,
	) (*discordgo.Member, error)
}

type db interface {
	CountUserQuotesInGuild(
		ctx context.Context, guildID models.ID, authorID models.ID,
	) (int, error)

	GetUserQuotesInGuild(
		ctx context.Context,
		guildID models.ID,
		authorID models.ID,
		offset, limit int,
	) ([]*models.Quote, error)
}
