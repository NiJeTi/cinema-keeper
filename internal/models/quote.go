package models

import (
	"time"
)

type Quote struct {
	AuthorID  DiscordID
	Text      string
	GuildID   DiscordID
	AddedByID DiscordID
	Timestamp time.Time
}
