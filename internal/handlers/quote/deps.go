package quote

import (
	"context"

	"github.com/nijeti/cinema-keeper/internal/models"
)

type db interface {
	GetUserQuotesOnGuild(
		ctx context.Context, authorID models.ID, guildID models.ID,
	) ([]*models.Quote, error)

	AddUserQuoteOnGuild(
		ctx context.Context, quote *models.Quote,
	) error
}
