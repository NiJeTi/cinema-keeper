package commands

import (
	"github.com/bwmarrin/discordgo"
)

const (
	QuoteName = "quote"

	QuoteSubCommandGet = "get"
	QuoteSubCommandAdd = "add"

	QuoteOptionAuthor = "author"
	QuoteOptionText   = "text"
)

const QuoteMaxQuotesPerPage = 10

func Quote() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        QuoteName,
		Description: "Collection of users' quotes",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        QuoteSubCommandGet,
				Description: "Get a random quote or a list of quotes from a user",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        QuoteOptionAuthor,
						Description: "List quotes from this user",
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        QuoteSubCommandAdd,
				Description: "Add a new quote",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        QuoteOptionAuthor,
						Description: "User who said that",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        QuoteOptionText,
						Description: "Text of the quote",
						Required:    true,
					},
				},
			},
		},
	}
}
