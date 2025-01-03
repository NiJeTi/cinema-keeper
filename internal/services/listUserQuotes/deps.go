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

	SendEmbeds(
		ctx context.Context,
		channelID models.ID,
		embeds []*discordgo.MessageEmbed,
	) error

	GuildMember(
		ctx context.Context, guildID models.ID, userID models.ID,
	) (*discordgo.Member, error)
}

type db interface {
	GetUserQuotesInGuild(
		ctx context.Context, guildID models.ID, authorID models.ID,
	) ([]*models.Quote, error)
}
