package commands

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/discordUtils"
)

const (
	QuoteName = "quote"

	QuoteSubCommandGet = "get"
	QuoteSubCommandAdd = "add"

	QuoteOptionAuthor = "author"
	QuoteOptionText   = "text"
)

const QuoteMaxQuotesPerPage = 5

const (
	quotesPageCustomIDAuthorIDIndex = 3
	quotesPageCustomIDPageIndex     = 4
)

var quotesPageCustomIDRegex = regexp.MustCompile(
	`^([a-z]+)_([a-z]+)_(\d+)_(\d+)$`,
)

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

func QuotesButtonCustomID(authorID models.ID, page int) string {
	if page < 0 {
		panic("page must be positive")
	}

	return fmt.Sprintf(
		"%s_%s_%s_%d", QuoteName, QuoteSubCommandGet, authorID, page,
	)
}

func QuoteParseButtonCustomID(id string) (authorID models.ID, page int) {
	matches := quotesPageCustomIDRegex.FindStringSubmatch(id)
	if len(matches) < quotesPageCustomIDRegex.NumSubexp() {
		panic(discordUtils.ErrInvalidCustomID)
	}

	authorID = models.ID(matches[quotesPageCustomIDAuthorIDIndex])

	page, err := strconv.Atoi(matches[quotesPageCustomIDPageIndex])
	if err != nil {
		panic(fmt.Errorf("failed to parse page: %w", err))
	}

	return authorID, page
}
