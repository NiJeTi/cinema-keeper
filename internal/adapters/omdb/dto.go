package omdb

import (
	"github.com/nijeti/cinema-keeper/internal/models"
)

//nolint:tagliatelle // external model
type movie struct {
	IMDBID string `json:"imdbID"`
	Title  string `json:"Title"`
}

type movieSearch struct {
	Search []movie
}

func (m movie) toModel() models.MovieShort {
	return models.MovieShort{
		ID:    models.IMDBID(m.IMDBID),
		Title: m.Title,
	}
}

func (s movieSearch) toModel() []models.MovieShort {
	movies := make([]models.MovieShort, 0, len(s.Search))
	for _, m := range s.Search {
		movies = append(movies, m.toModel())
	}
	return movies
}
