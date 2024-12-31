package diceRoll

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type discord interface {
	Respond(
		ctx context.Context,
		i *discordgo.Interaction,
		response *discordgo.InteractionResponse,
	) error
}
