package models

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type Quote struct {
	Timestamp time.Time
	Author    discordgo.Member
	AddedBy   discordgo.Member
	GuildID   ID
	Text      string
}
