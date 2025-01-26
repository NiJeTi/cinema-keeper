package addQuote

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
		ctx context.Context, guildID models.DiscordID, userID models.DiscordID,
	) (*discordgo.Member, error)
}

type db interface {
	AddUserQuoteInGuild(ctx context.Context, quote *models.Quote) error
}
