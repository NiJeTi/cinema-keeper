package searchNewMovie_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands"
	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/searchNewMovie"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/searchNewMovie"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{}

	type setup func(
		t *testing.T, title string, err error,
	) *searchNewMovie.Service

	tests := map[string]struct {
		title string
		err   error
		setup setup
	}{
		"empty_title_response_error": {
			title: "",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, _ string, err error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(err)

				return searchNewMovie.New(d, omdb)
			},
		},

		"omdb_error": {
			title: "title",
			err:   errors.New("omdb error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				omdb.EXPECT().MoviesByTitle(ctx, title).Return(nil, err)

				return searchNewMovie.New(d, omdb)
			},
		},
		"no_movies_response_error": {
			title: "title",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				omdb.EXPECT().MoviesByTitle(
					ctx, title,
				).Return([]models.MovieBase{}, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(err)

				return searchNewMovie.New(d, omdb)
			},
		},
		"movies_response_error": {
			title: "title",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				movies := []models.MovieBase{{ID: "id", Title: title}}
				omdb.EXPECT().MoviesByTitle(ctx, title).Return(movies, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteSearch(movies),
				).Return(err)

				return searchNewMovie.New(d, omdb)
			},
		},
		"empty_title": {
			title: "",
			err:   nil,
			setup: func(
				t *testing.T, _ string, _ error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(nil)

				return searchNewMovie.New(d, omdb)
			},
		},
		"no_movies": {
			title: "title",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				omdb.EXPECT().MoviesByTitle(
					ctx, title,
				).Return([]models.MovieBase{}, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteEmptySearch(),
				).Return(nil)

				return searchNewMovie.New(d, omdb)
			},
		},
		"too_many_movies": {
			title: "title",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				var movies []models.MovieBase
				for i := range commands.MovieOptionTitleChoiceLimit + 1 {
					movies = append(
						movies, models.MovieBase{
							ID:    models.ImdbID(fmt.Sprint(i)),
							Title: title,
						},
					)
				}
				omdb.EXPECT().MoviesByTitle(ctx, title).Return(movies, nil)

				resp := responses.MovieAutocompleteSearch(
					movies[:commands.MovieOptionTitleChoiceLimit],
				)
				d.EXPECT().Respond(
					ctx, i, resp,
				).Return(nil)

				return searchNewMovie.New(d, omdb)
			},
		},
		"success": {
			title: "title",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *searchNewMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)

				var movies []models.MovieBase
				for i := range commands.MovieOptionTitleChoiceLimit {
					movies = append(
						movies, models.MovieBase{
							ID:    models.ImdbID(fmt.Sprint(i)),
							Title: title,
						},
					)
				}
				omdb.EXPECT().MoviesByTitle(ctx, title).Return(movies, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAutocompleteSearch(movies),
				).Return(nil)

				return searchNewMovie.New(d, omdb)
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
