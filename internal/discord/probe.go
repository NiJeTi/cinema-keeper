package discord

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type Probe struct {
	session *discordgo.Session
}

func NewProbe(
	session *discordgo.Session,
) *Probe {
	return &Probe{
		session: session,
	}
}

func (p *Probe) Check(_ context.Context) error {
	_, err := p.session.User("@me")
	return err //nolint:wrapcheck // error passthrough
}
