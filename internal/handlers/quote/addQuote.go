package quote

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
)

func (h *Handler) addQuote(
	i *discordgo.InteractionCreate,
	author discordgo.Member,
	text string,
) {
	quote := &models.Quote{
		Author:  author,
		Text:    text,
		GuildID: models.ID(i.GuildID),
		AddedBy: *i.Member,
	}
	err := h.db.AddUserQuoteOnGuild(h.ctx, quote)
	if err != nil {
		return
	}

	quote.Author = author
	quote.AddedBy = *i.Interaction.Member

	_ = h.utils.Respond(i, responses.QuoteAdded(quote))
}
