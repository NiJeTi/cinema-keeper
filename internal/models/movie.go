package models

import (
	"time"
)

type MovieBase struct {
	ID    IMDBID
	Title string
	Year  int
}

type MovieMeta struct {
	MovieBase
	Director  string
	Plot      string
	PosterURL string
}

type Movie struct {
	MovieBase
	AddedByID ID
	AddedAt   time.Time
	WatchedAt *time.Time
	Rating    *int
}
