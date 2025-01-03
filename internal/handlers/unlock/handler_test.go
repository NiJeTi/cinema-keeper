package unlock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"

	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/handlers/unlock"
	"github.com/nijeti/cinema-keeper/internal/handlers/unlock"
)

func TestHandler_Handle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	i := &discordgo.InteractionCreate{}

	type setup func(t *testing.T, err error) *unlock.Handler

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"service_error": {
			err: errors.New("service error"),
			setup: func(t *testing.T, err error) *unlock.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(ctx, i.Interaction).Return(err)

				return unlock.New(s)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *unlock.Handler {
				s := mocks.NewMockService(t)

				s.EXPECT().Exec(ctx, i.Interaction).Return(nil)

				return unlock.New(s)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				h := tt.setup(t, tt.err)
				err := h.Handle(ctx, i)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
