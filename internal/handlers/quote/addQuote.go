package quote

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/types"
)

func (h Handler) addQuote(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	author *discordgo.Member,
	text string,
) {
	quote := &models.Quote{
		AuthorID:  types.ID(author.User.ID),
		Text:      text,
		GuildID:   types.ID(i.GuildID),
		AddedByID: types.ID(i.Member.User.ID),
	}
	err := h.db.AddUserQuoteOnGuild(h.ctx, quote)
	if err != nil {
		return
	}

	quote.Author = author
	quote.AddedBy = i.Interaction.Member

	_ = h.utils.Respond(h.ctx, s, i, responses.QuoteAdded(quote))
}
