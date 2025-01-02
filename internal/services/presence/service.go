package presence

import (
	"context"
	"fmt"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
)

type Service struct {
	discord discord
}

func New(
	discord discord,
) *Service {
	return &Service{
		discord: discord,
	}
}

func (s *Service) Set(ctx context.Context) error {
	err := s.discord.SetActivity(ctx, responses.Activity())
	if err != nil {
		return fmt.Errorf("failed to set activity: %w", err)
	}

	return nil
}
