package models

import (
	"time"
)

type Quote struct {
	AuthorID  ID
	Text      string
	GuildID   ID
	AddedByID ID
	Timestamp time.Time
}
