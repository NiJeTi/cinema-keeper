package addMovie_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/addMovie"
	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/services/addMovie"
)

func TestService_Exec(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.Interaction{
		GuildID: "1",
		Member: &discordgo.Member{
			User: &discordgo.User{
				ID: "2",
			},
		},
	}

	type setup func(t *testing.T, title string, err error) *addMovie.Service

	tests := map[string]struct {
		title string
		err   error
		setup setup
	}{
		"invalid_title_response_error": {
			title: "invalid",
			err:   errors.New("response error"),
			setup: func(t *testing.T, _ string, err error) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieInvalidTitle(),
				).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"movie_by_imdb_id_error": {
			title: "tt123",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				db.EXPECT().MovieByImdbID(
					ctx, models.ImdbID(title),
				).Return(nil, nil, err)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie_guild_contains_db_error": {
			title: "tt123",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				movieID := models.ID(123)
				movie := models.MovieMeta{}
				db.EXPECT().MovieByImdbID(
					ctx, models.ImdbID(title),
				).Return(&movieID, &movie, nil)

				db.EXPECT().GuildMovieExists(
					ctx, movieID, models.DiscordID(i.GuildID),
				).Return(false, err)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie_guild_contains_response_error": {
			title: "tt123",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				movieID := models.ID(123)
				movie := models.MovieMeta{}
				db.EXPECT().MovieByImdbID(
					ctx, models.ImdbID(title),
				).Return(&movieID, &movie, nil)

				db.EXPECT().GuildMovieExists(
					ctx, movieID, models.DiscordID(i.GuildID),
				).Return(true, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAlreadyExists(),
				).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"new_movie_omdb_error": {
			title: "tt123",
			err:   errors.New("omdb error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)

				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(nil, nil, nil)

				omdb.EXPECT().MovieByID(ctx, id).Return(nil, err)

				return addMovie.New(d, omdb, db)
			},
		},
		"new_movie_db_error": {
			title: "tt123",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)

				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(nil, nil, nil)

				movieMeta := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				omdb.EXPECT().MovieByID(ctx, id).Return(movieMeta, nil)

				db.EXPECT().AddMovie(ctx, movieMeta).Return(0, err)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie_add_to_guild_error": {
			title: "tt123",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				movieID := models.ID(123)
				movie := models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(&movieID, &movie, nil)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMovieExists(
					ctx, movieID, guildID,
				).Return(false, nil)

				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie_response_error": {
			title: "tt123",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				movieID := models.ID(123)
				movie := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(&movieID, movie, nil)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMovieExists(
					ctx, movieID, guildID,
				).Return(false, nil)

				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"new_movie_add_to_guild_error": {
			title: "tt123",
			err:   errors.New("db error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(nil, nil, nil)

				movieMeta := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				omdb.EXPECT().MovieByID(ctx, id).Return(movieMeta, nil)

				movieID := models.ID(123)
				db.EXPECT().AddMovie(ctx, movieMeta).Return(movieID, nil)

				guildID := models.DiscordID(i.GuildID)
				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"new_movie_response_error": {
			title: "tt123",
			err:   errors.New("response error"),
			setup: func(
				t *testing.T, title string, err error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(nil, nil, nil)

				movieMeta := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				omdb.EXPECT().MovieByID(ctx, id).Return(movieMeta, nil)

				movieID := models.ID(123)
				db.EXPECT().AddMovie(ctx, movieMeta).Return(movieID, nil)

				guildID := models.DiscordID(i.GuildID)
				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(err)

				return addMovie.New(d, omdb, db)
			},
		},
		"invalid_title": {
			title: "invalid",
			err:   nil,
			setup: func(t *testing.T, _ string, _ error) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				d.EXPECT().Respond(
					ctx, i, responses.MovieInvalidTitle(),
				).Return(nil)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie": {
			title: "tt123",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				movieID := models.ID(123)
				movie := models.MovieMeta{}
				db.EXPECT().MovieByImdbID(
					ctx, models.ImdbID(title),
				).Return(&movieID, &movie, nil)

				db.EXPECT().GuildMovieExists(
					ctx, movieID, models.DiscordID(i.GuildID),
				).Return(true, nil)

				d.EXPECT().Respond(
					ctx, i, responses.MovieAlreadyExists(),
				).Return(nil)

				return addMovie.New(d, omdb, db)
			},
		},
		"existing_movie_success": {
			title: "tt123",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				movieID := models.ID(123)
				movie := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(&movieID, movie, nil)

				guildID := models.DiscordID(i.GuildID)
				db.EXPECT().GuildMovieExists(
					ctx, movieID, guildID,
				).Return(false, nil)

				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(nil)

				return addMovie.New(d, omdb, db)
			},
		},
		"new_movie_success": {
			title: "tt123",
			err:   nil,
			setup: func(
				t *testing.T, title string, _ error,
			) *addMovie.Service {
				d := mocks.NewMockDiscord(t)
				omdb := mocks.NewMockOmdb(t)
				db := mocks.NewMockDb(t)

				id := models.ImdbID(title)
				db.EXPECT().MovieByImdbID(
					ctx, id,
				).Return(nil, nil, nil)

				movieMeta := &models.MovieMeta{
					MovieBase: models.MovieBase{
						ID: id,
					},
				}
				omdb.EXPECT().MovieByID(ctx, id).Return(movieMeta, nil)

				movieID := models.ID(123)
				db.EXPECT().AddMovie(ctx, movieMeta).Return(movieID, nil)

				guildID := models.DiscordID(i.GuildID)
				addedByID := models.DiscordID(i.Member.User.ID)
				db.EXPECT().AddMovieToGuild(
					ctx, movieID, guildID, addedByID,
				).Return(nil)

				d.EXPECT().Respond(ctx, i, mock.Anything).Return(nil)

				return addMovie.New(d, omdb, db)
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
