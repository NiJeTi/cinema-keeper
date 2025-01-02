package presence

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type discord interface {
	SetActivity(ctx context.Context, activity *discordgo.Activity) error
}
