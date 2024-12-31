package roll

import (
	"context"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/dice"
)

type service interface {
	Exec(
		ctx context.Context, i *discordgo.Interaction, size dice.Size,
		count int,
	) error
}
