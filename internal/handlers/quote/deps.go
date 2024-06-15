package quote

import (
	"context"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/types"
)

type db interface {
	GetUserQuotesOnGuild(
		ctx context.Context,
		authorID types.ID,
		guildID types.ID,
	) ([]*models.Quote, error)

	AddUserQuoteOnGuild(
		ctx context.Context,
		quote *models.Quote,
	) error
}
