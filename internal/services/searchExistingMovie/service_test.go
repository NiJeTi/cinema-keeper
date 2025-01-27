package searchExistingMovie_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/searchExistingMovie"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/searchExistingMovie"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
	}

	type setup func(
		t *testing.T, title string, err error,
	) *searchExistingMovie.Service

	tests := map[string]struct {
		title string
		err   error
		setup setup
	}{
		"title_empty_response_error": {
			title: "",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, _ string, err error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(err)

				return searchExistingMovie.New(d, db)
			},
		},
		"db_error": {
			title: "title",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMoviesByTitle(
					ctx, guildID, title,
					commands.MovieTitleAutocompleteChoicesLimit,
				).Return(nil, err)

				return searchExistingMovie.New(d, db)
			},
		},
		"no_movies_response_error": {
			title: "title",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMoviesByTitle(
					ctx, guildID, title,
					commands.MovieTitleAutocompleteChoicesLimit,
				).Return([]models.MovieBase{}, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(err)

				return searchExistingMovie.New(d, db)
			},
		},
		"response_error": {
			title: "title",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				guildID := models.DiscordID(i.GuildID)
				movies := []models.MovieBase{
					{ID: "id", Title: title},
				}
				db.EXPECT().GuildMoviesByTitle(
					ctx, guildID, title,
					commands.MovieTitleAutocompleteChoicesLimit,
				).Return(movies, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteSearch(movies),
				).Return(err)

				return searchExistingMovie.New(d, db)
			},
		},
		"title_empty": {
			title: "",
			err:   nil,
			setup: func(
				t *testing.T, _ string, _ error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(nil)

				return searchExistingMovie.New(d, db)
			},
		},
		"no_movies": {
			title: "title",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMoviesByTitle(
					ctx, guildID, title,
					commands.MovieTitleAutocompleteChoicesLimit,
				).Return([]models.MovieBase{}, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(nil)

				return searchExistingMovie.New(d, db)
			},
		},
		"success": {
			title: "title",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *searchExistingMovie.Service {
				d := mocks.NewMockDiscord(t)
				db := mocks.NewMockDb(t)

				guildID := models.DiscordID(i.GuildID)
				movies := []models.MovieBase{
					{ID: "id", Title: title},
				}
				db.EXPECT().GuildMoviesByTitle(
					ctx, guildID, title,
					commands.MovieTitleAutocompleteChoicesLimit,
				).Return(movies, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteSearch(movies),
				).Return(nil)

				return searchExistingMovie.New(d, db)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.title, tt.err)
				err := s.Exec(ctx, i, tt.title)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
