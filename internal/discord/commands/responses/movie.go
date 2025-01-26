package responses

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
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
				Name:  movie.Title,
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
