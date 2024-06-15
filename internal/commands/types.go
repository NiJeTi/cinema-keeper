package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type Command struct {
	Description *discordgo.ApplicationCommand
	Handler     Handler
}
