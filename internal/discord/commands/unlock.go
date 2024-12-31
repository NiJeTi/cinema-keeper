package commands

import (
	"github.com/bwmarrin/discordgo"
)

const (
	UnlockName = "unlock"
)

func Unlock() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        UnlockName,
		Description: "Remove user limit for current voice channel",
	}
}
