package responses

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/utils"
)

func QuoteGuildNoQuotes() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "There are no quotes in this server",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func QuoteUserNoQuotes(
	author *discordgo.Member,
) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"%s doesn't have any quotes", author.Mention(),
			),
		},
	}
}

func QuoteRandomQuote(quote *models.Quote) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:     quote.Text,
					Timestamp: quote.Timestamp.Format(time.RFC3339),
					Color:     utils.RandomColor(),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    quote.Author.DisplayName(),
						IconURL: quote.Author.AvatarURL("16x16"),
					},
				},
			},
		},
	}
}

func QuoteList(
	quotes []*models.Quote, page int, lastPage int,
) *discordgo.InteractionResponse {
	if len(quotes) > commands.QuoteMaxQuotesPerPage {
		panic("too many quotes per page")
	}

	author := quotes[0].Author

	prevButton := discordgo.Button{
		Style: discordgo.PrimaryButton,
		Emoji: &discordgo.ComponentEmoji{
			Name: "◀️",
		},
	}
	if page == 0 {
		prevButton.Disabled = true
		prevButton.CustomID = "prev"
	} else {
		prevButton.Disabled = false
		prevButton.CustomID = commands.QuotesButtonCustomID(
			models.ID(author.User.ID), page-1,
		)
	}

	nextButton := discordgo.Button{
		Style: discordgo.PrimaryButton,
		Emoji: &discordgo.ComponentEmoji{
			Name: "▶️",
		},
	}
	if page == lastPage {
		nextButton.Disabled = true
		nextButton.CustomID = "next"
	} else {
		nextButton.Disabled = false
		nextButton.CustomID = commands.QuotesButtonCustomID(
			models.ID(author.User.ID), page+1,
		)
	}

	headerEmbed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(
			"The most stunning quotes of %s", author.Mention(),
		),
		Color: utils.RandomColor(),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: author.AvatarURL("64x64"),
		},
	}

	quoteEmbeds := make([]*discordgo.MessageEmbed, 0, len(quotes))
	for _, quote := range quotes {
		embed := &discordgo.MessageEmbed{
			Title:     quote.Text,
			Timestamp: quote.Timestamp.Format(time.RFC3339),
			Color:     utils.RandomColor(),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    quote.AddedBy.DisplayName(),
				IconURL: quote.AddedBy.AvatarURL("16x16"),
			},
		}
		quoteEmbeds = append(quoteEmbeds, embed)
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: append(
				[]*discordgo.MessageEmbed{headerEmbed}, quoteEmbeds...,
			),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						prevButton, nextButton,
					},
				},
			},
		},
	}
}

func QuoteAdded(quote *models.Quote) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Quote added",
					Description: quote.Text,
					Timestamp:   quote.Timestamp.Format(time.RFC3339),
					Color:       utils.RandomColor(),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    quote.Author.DisplayName(),
						IconURL: quote.Author.AvatarURL("16x16"),
					},
				},
			},
		},
	}
}
