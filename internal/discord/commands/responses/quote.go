package responses

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

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
					Author: &discordgo.MessageEmbedAuthor{
						Name:    quote.Author.DisplayName(),
						IconURL: quote.Author.AvatarURL("32x32"),
					},
				},
			},
		},
	}
}

func QuoteListHeader(author *discordgo.Member) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Quotes",
					Description: fmt.Sprintf(
						"The most stunning quotes of %s", author.Mention(),
					),
					Color: utils.RandomColor(),
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: author.AvatarURL("64x64"),
					},
				},
			},
		},
	}
}

func QuoteList(quotes []*models.Quote) []*discordgo.MessageEmbed {
	embeds := make([]*discordgo.MessageEmbed, 0, len(quotes))
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
		embeds = append(embeds, embed)
	}
	return embeds
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
					Author: &discordgo.MessageEmbedAuthor{
						Name:    quote.AddedBy.DisplayName(),
						IconURL: quote.AddedBy.AvatarURL("32x32"),
					},
				},
			},
		},
	}
}
