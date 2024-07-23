package discord

import (
	"github.com/bwmarrin/discordgo"
)

const (
	UnlockName = "unlock"
)

var Unlock = &discordgo.ApplicationCommand{
	Name:        UnlockName,
	Description: "Remove user limit for current voice channel",
}
