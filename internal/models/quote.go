package models

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Quote struct {
	Author    *discordgo.Member
	Text      string
	GuildID   ID
	AddedBy   *discordgo.Member
	Timestamp time.Time
}
