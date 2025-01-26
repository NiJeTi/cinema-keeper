//nolint:tagliatelle // external API DTOs
package omdb

import (
	"github.com/nijeti/cinema-keeper/internal/models"
)

type movieBase struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   string `json:"Year"`
}

type movie struct {
	movieBase
	Genre     string `json:"Genre"`
	Director  string `json:"Director"`
	Plot      string `json:"Plot"`
	PosterURL string `json:"Poster"`
}

type search struct {
	Search []movieBase `json:"Search"`
}

func (m movieBase) toModel() models.MovieBase {
	return models.MovieBase{
		ID:    models.ImdbID(m.ImdbID),
		Title: m.Title,
		Year:  m.Year,
	}
}

func (m *movie) toModel() *models.MovieMeta {
	return &models.MovieMeta{
		MovieBase: m.movieBase.toModel(),
		Genre:     m.Genre,
		Director:  m.Director,
		Plot:      m.Plot,
		PosterURL: m.PosterURL,
	}
}

func (s search) toModel() []models.MovieBase {
	movies := make([]models.MovieBase, 0, len(s.Search))
	for _, m := range s.Search {
		movies = append(movies, m.toModel())
	}
	return movies
}
