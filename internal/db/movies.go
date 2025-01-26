package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

type MoviesRepo struct {
	repoBase
}

type movieRow struct {
	ID        int64  `db:"id"`
	ImdbID    string `db:"imdb_id"`
	Title     string `db:"title"`
	Year      string `db:"year"`
	Genre     string `db:"genre"`
	Director  string `db:"director"`
	Plot      string `db:"plot"`
	PosterURL string `db:"poster_url"`
}

func NewMoviesRepo(
	db *sqlx.DB,
) *MoviesRepo {
	return &MoviesRepo{
		repoBase: repoBase{
			db: db,
		},
	}
}

func (r *MoviesRepo) MovieByImdbID(
	ctx context.Context, id models.ImdbID,
) (*models.ID, *models.MovieMeta, error) {
	const query = `
		select * from movies
		where imdb_id = $1`

	var row movieRow
	err := r.db.GetContext(ctx, &row, query, id)
	switch {
	case err == nil:
		return ptr.To(models.ID(row.ID)), row.toMovieMeta(), nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil, nil
	default:
		return nil, nil, fmt.Errorf("failed to query movie: %w", err)
	}
}

func (r *MoviesRepo) AddMovie(
	ctx context.Context, movie *models.MovieMeta,
) (models.ID, error) {
	const query = `
		insert into movies(imdb_id, title, year, genre, director, plot, poster_url)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning id`

	var movieID int64
	return models.ID(movieID), r.withTransaction(
		ctx, func(ctx context.Context, tx *sqlx.Tx) error {
			err := tx.QueryRowContext(
				ctx,
				query,
				movie.ID,
				movie.Title,
				movie.Year,
				movie.Genre,
				movie.Director,
				movie.Plot,
				movie.PosterURL,
			).Scan(&movieID)
			if err != nil {
				return fmt.Errorf("failed to insert movie: %w", err)
			}

			return nil
		},
	)
}

func (r *MoviesRepo) GuildMovieExists(
	ctx context.Context, movieID models.ID, guildID models.DiscordID,
) (bool, error) {
	const query = `select exists(
		select 1 from guild_movies
		where movie_id = $1 and guild_id = $2)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, movieID, guildID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf(
			"failed to check if guild movie exists: %w", err,
		)
	}

	return exists, nil
}

func (r *MoviesRepo) AddMovieToGuild(
	ctx context.Context,
	movieID models.ID,
	guildID models.DiscordID,
	addedByID models.DiscordID,
) error {
	const query = `
		insert into guild_movies(movie_id, guild_id, added_by_id, added_at)
		values ($1, $2, $3, $4)`

	return r.withTransaction(
		ctx, func(ctx context.Context, tx *sqlx.Tx) error {
			_, err := tx.ExecContext(
				ctx, query, movieID, guildID, addedByID, time.Now().UTC(),
			)
			if err != nil {
				return fmt.Errorf("failed to insert guild movie: %w", err)
			}
			return nil
		},
	)
}

func (r *movieRow) toMovieBase() models.MovieBase {
	return models.MovieBase{
		ID:    models.ImdbID(r.ImdbID),
		Title: r.Title,
		Year:  r.Year,
	}
}

func (r *movieRow) toMovieMeta() *models.MovieMeta {
	return &models.MovieMeta{
		MovieBase: r.toMovieBase(),
		Genre:     r.Director,
		Director:  r.Director,
		Plot:      r.Plot,
		PosterURL: r.PosterURL,
	}
}
