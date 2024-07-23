package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type Probe struct {
	discord *discordgo.Session
}

func NewProbe(
	discord *discordgo.Session,
) Probe {
	return Probe{
		discord: discord,
	}
}

func (p Probe) Check(_ context.Context) error {
	_, err := p.discord.User("@me")
	return err
}
