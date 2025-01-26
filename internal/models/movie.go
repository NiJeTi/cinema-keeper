package models

import (
	"time"
)

type MovieShort struct {
	ID    IMDBID
	Title string
}

type Movie struct {
	MovieShort
	AddedByID ID
	AddedAt   time.Time
	WatchedAt *time.Time
	Rating    *int
}
