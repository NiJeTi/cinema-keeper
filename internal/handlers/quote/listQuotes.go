package quote

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/responses"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/types"
)

func (h Handler) listQuotes(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	author *discordgo.Member,
) {
	quotes, err := h.db.GetUserQuotesOnGuild(
		h.ctx,
		types.ID(author.User.ID),
		types.ID(i.GuildID),
	)
	if err != nil {
		return
	}

	if len(quotes) == 0 {
		_ = h.utils.Respond(h.ctx, s, i, responses.QuotesEmpty(author))
		return
	}

	err = h.enrichQuotes(s, i, quotes, author)
	if err != nil {
		return
	}

	err = h.utils.Respond(h.ctx, s, i, responses.QuotesHeader(author))
	if err != nil {
		return
	}

	for index := 0; index < len(quotes); index += 10 {
		limit := index + 10
		if index+10 >= len(quotes) {
			limit = len(quotes)
		}

		err = h.sendEmbeds(s, i, responses.Quotes(quotes[index:limit]))
		if err != nil {
			return
		}
	}
}

func (h Handler) enrichQuotes(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	quotes []*models.Quote,
	author *discordgo.Member,
) error {
	for _, q := range quotes {
		q.Author = author

		addedBy, err := s.State.Member(i.GuildID, q.AddedByID.String())
		if err != nil {
			h.log.ErrorContext(h.ctx, "failed to get member", "error", err)
			return err
		}
		q.AddedBy = addedBy
	}
	return nil
}

func (h Handler) sendEmbeds(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	embeds []*discordgo.MessageEmbed,
) error {
	_, err := s.ChannelMessageSendEmbeds(i.Interaction.ChannelID, embeds)
	if err != nil {
		h.log.ErrorContext(h.ctx, "failed to send embeds", "error", err)
	}
	return err
}
