//nolint:tagliatelle // external API DTOs
package omdb

import (
	"fmt"
	"strconv"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type movieBase struct {
	IMDBID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   string `json:"Year"`
}

type movie struct {
	movieBase
	Director  string `json:"Director"`
	Plot      string `json:"Plot"`
	PosterURL string `json:"Poster"`
}

type search struct {
	Search []movieBase `json:"Search"`
}

func (m movieBase) toModel() models.MovieBase {
	year, err := strconv.Atoi(m.Year)
	if err != nil {
		panic(fmt.Errorf("failed to parse year: %w", err))
	}

	return models.MovieBase{
		ID:    models.IMDBID(m.IMDBID),
		Title: m.Title,
		Year:  year,
	}
}

func (m *movie) toModel() *models.MovieMeta {
	return &models.MovieMeta{
		MovieBase: m.movieBase.toModel(),
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
