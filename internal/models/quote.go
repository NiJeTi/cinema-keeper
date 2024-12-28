package models

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type Quote struct {
	Author    *discordgo.Member
	AddedBy   *discordgo.Member
	Text      string
	Timestamp time.Time
	AuthorID  types.ID
	AddedByID types.ID
	GuildID   types.ID
}
