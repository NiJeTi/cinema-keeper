package responses

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/die"
	"github.com/nijeti/cinema-keeper/internal/pkg/utils"
)

func RollSingleDefault(result int) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strconv.Itoa(result),
		},
	}
}

func RollSingle(size die.Size, result int) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**D%d**: %d", size, result),
		},
	}
}

func RollMultiple(size die.Size, results []int) *discordgo.InteractionResponse {
	var text string
	sum := 0
	for _, result := range results {
		text += fmt.Sprintf("> %d\n", result)
		sum += result
	}
	text += fmt.Sprintf("\n**Sum**: %d", sum)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("D%d", size),
					Description: text,
					Color:       utils.RandomColor(),
				},
			},
		},
	}
}
