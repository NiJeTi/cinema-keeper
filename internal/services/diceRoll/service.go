package diceRoll

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
)

type Service struct {
	discord discord
}

func New(
	discord discord,
) *Service {
	return &Service{
		discord: discord,
	}
}

func (s *Service) Exec(
	ctx context.Context, i *discordgo.Interaction, size dice.Size, count int,
) error {
	rolls := make([]int, 0, count)
	for range count {
		r := dice.Roll(size)
		rolls = append(rolls, r)
	}

	var response *discordgo.InteractionResponse
	if len(rolls) == commands.RollOptionCountDefault {
		if size == commands.RollOptionSizeDefault {
			response = responses.RollSingleDefault(rolls[0])
		} else {
			response = responses.RollSingle(size, rolls[0])
		}
	} else {
		response = responses.RollMultiple(size, rolls)
	}

	err := s.discord.Respond(ctx, i, response)
	if err != nil {
		return fmt.Errorf("failed to respond: %w", err)
	}

	return nil
}
