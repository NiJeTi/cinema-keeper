package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

const (
	MovieName = "movie"

	MovieSubCommandGet     = "get"
	MovieSubCommandAdd     = "add"
	MovieSubCommandRemove  = "remove"
	MovieSubCommandList    = "list"
	MovieSubCommandWatched = "watched"

	MovieOptionTitle = "title"

	MovieOptionType       = "type"
	MovieOptionTypeNext   = "next"
	MovieOptionTypeRandom = "random"

	MovieOptionRate        = "rate"
	MovieOptionRateLike    = "like"
	MovieOptionRateDislike = "dislike"
)

const (
	MovieOptionTypeDefault = movieOptionTypeNextValue
	MovieOptionRateDefault = 0

	MovieOptionTitleMinLength   = 4
	MovieOptionTitleChoiceLimit = 5
)

const (
	movieOptionTypeNextValue   = 0
	movieOptionTypeRandomValue = 1

	movieOptionRateLikeValue    = 1
	movieOptionRateDislikeValue = -1
)

const (
	movieOptionTitleValueMinLength = 9
	movieOptionTitleValueMaxLength = 9
)

func Movie() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        MovieName,
		Description: "Manage movie list",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        MovieSubCommandGet,
				Description: "Get a random movie or a next one in the queue",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        MovieOptionType,
						Description: "Type of movie to get",
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  MovieOptionTypeNext,
								Value: movieOptionTypeNextValue,
							},
							{
								Name:  MovieOptionTypeRandom,
								Value: movieOptionTypeRandomValue,
							},
						},
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        MovieSubCommandAdd,
				Description: "Add a movie to the queue",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:         discordgo.ApplicationCommandOptionString,
						Name:         MovieOptionTitle,
						Description:  "Title of the movie",
						Required:     true,
						Autocomplete: true,
						MinLength:    ptr.To(movieOptionTitleValueMinLength),
						MaxLength:    movieOptionTitleValueMaxLength,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        MovieSubCommandRemove,
				Description: "Remove a movie from the queue",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:         discordgo.ApplicationCommandOptionString,
						Name:         MovieOptionTitle,
						Description:  "Title of the movie",
						Required:     true,
						Autocomplete: true,
						MinLength:    ptr.To(movieOptionTitleValueMinLength),
						MaxLength:    movieOptionTitleValueMaxLength,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        MovieSubCommandList,
				Description: "List all movies in the queue",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        MovieSubCommandWatched,
				Description: "Mark a movie as watched",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:         discordgo.ApplicationCommandOptionString,
						Name:         MovieOptionTitle,
						Description:  "Title of the movie",
						Required:     true,
						Autocomplete: true,
						MinLength:    ptr.To(movieOptionTitleValueMinLength),
						MaxLength:    movieOptionTitleValueMaxLength,
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        MovieOptionRate,
						Description: "Rate the movie",
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  MovieOptionRateLike,
								Value: movieOptionRateLikeValue,
							},
							{
								Name:  MovieOptionRateDislike,
								Value: movieOptionRateDislikeValue,
							},
						},
					},
				},
			},
		},
	}
}
