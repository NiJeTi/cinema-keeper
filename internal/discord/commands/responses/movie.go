package responses

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/utils"
)

func MovieAutocompleteEmptySearch() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: []*discordgo.ApplicationCommandOptionChoice{},
		},
	}
}

func MovieAutocompleteSearch(
	movies []models.MovieBase,
) *discordgo.InteractionResponse {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(movies))
	for _, movie := range movies {
		choices = append(
			choices, &discordgo.ApplicationCommandOptionChoice{
				Name:  fmt.Sprintf("%s (%s)", movie.Title, movie.Year),
				Value: movie.ID,
			},
		)
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	}
}

func MovieInvalidTitle() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Invalid title. Please select from the list.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func MovieAlreadyExists() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Movie is already in list",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}
}

func MovieAdded(movie *models.MovieMeta) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       movie.Title,
					Description: movie.Plot,
					Color:       utils.RandomColor(),
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: movie.PosterURL,
					},
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  "Year",
							Value: movie.Year,
						},
						{
							Name:  "Genre",
							Value: movie.Genre,
						},
						{
							Name:  "Director",
							Value: movie.Director,
						},
					},
				},
			},
		},
	}
}
