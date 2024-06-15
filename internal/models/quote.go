package models

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/types"
)

type Quote struct {
	Author    *discordgo.Member
	AuthorID  types.ID
	Text      string
	GuildID   types.ID
	AddedBy   *discordgo.Member
	AddedByID types.ID
	Timestamp time.Time
}
