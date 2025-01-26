package models

import (
	"time"
)

type MovieBase struct {
	ID    ImdbID
	Title string
	Year  string
}

type MovieMeta struct {
	MovieBase
	Genre     string
	Director  string
	Plot      string
	PosterURL string
}

type GuildMovie struct {
	MovieMeta
	GuildID   DiscordID
	AddedByID DiscordID
	AddedAt   time.Time
	WatchedAt *time.Time
	Rating    *MovieRating
}

type MovieRating int16

const (
	MovieRatingDislike MovieRating = -1
	MovieRatingLike    MovieRating = 1
)
