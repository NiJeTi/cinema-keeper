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

func QuoteRandomQuote(
	author *discordgo.Member, quote *models.Quote,
) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:     quote.Text,
					Timestamp: quote.Timestamp.Format(time.RFC3339),
					Color:     utils.RandomColor(),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    author.DisplayName(),
						IconURL: author.AvatarURL("16x16"),
					},
				},
			},
		},
	}
}

func QuoteList(
	author *discordgo.Member, quotes []*models.Quote, page, lastPage int,
) *discordgo.InteractionResponse {
	if len(quotes) > commands.QuoteMaxQuotesPerPage {
		panic("too many quotes per page")
	}

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
		Title:       "Quotes",
		Description: author.Mention(),
		Color:       utils.RandomColor(),
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

func QuoteAdded(
	author *discordgo.Member, quote *models.Quote,
) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "New quote",
					Description: quote.Text,
					Timestamp:   quote.Timestamp.Format(time.RFC3339),
					Color:       utils.RandomColor(),
					Footer: &discordgo.MessageEmbedFooter{
						Text:    author.DisplayName(),
						IconURL: author.AvatarURL("16x16"),
					},
				},
			},
		},
	}
}
