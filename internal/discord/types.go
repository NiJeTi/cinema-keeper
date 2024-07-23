package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Handler interface {
	Handle(i *discordgo.InteractionCreate)
}

type Command struct {
	Description *discordgo.ApplicationCommand
	Handler     Handler
}
