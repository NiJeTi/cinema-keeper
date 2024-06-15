package commands

import (
	"github.com/bwmarrin/discordgo"
)

const (
	QuoteName         = "quote"
	QuoteOptionAuthor = "author"
	QuoteOptionText   = "text"
)

var Quote = &discordgo.ApplicationCommand{
	Name:        QuoteName,
	Description: "Manage the mos stunning quotes of the specified user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        QuoteOptionAuthor,
			Description: "The user who said that",
			Type:        discordgo.ApplicationCommandOptionUser,
			Required:    true,
		},
		{
			Name:        QuoteOptionText,
			Description: "The text of the quote",
			Type:        discordgo.ApplicationCommandOptionString,
		},
	},
}
