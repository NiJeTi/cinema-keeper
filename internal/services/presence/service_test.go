package presence_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/discord/commands/responses"
	mocks "github.com/nijeti/cinema-keeper/internal/generated/mocks/services/presence"
	"github.com/nijeti/cinema-keeper/internal/services/presence"
)

func TestService_Set(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type setup func(t *testing.T, err error) *presence.Service

	tests := map[string]struct {
		err   error
		setup setup
	}{
		"set_activity_error": {
			err: errors.New("set activity error"),
			setup: func(t *testing.T, err error) *presence.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().SetActivity(ctx, responses.Activity()).Return(err)

				return presence.New(d)
			},
		},
		"success": {
			err: nil,
			setup: func(t *testing.T, _ error) *presence.Service {
				d := mocks.NewMockDiscord(t)

				d.EXPECT().SetActivity(ctx, responses.Activity()).Return(nil)

				return presence.New(d)
			},
		},
	}

	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				s := tt.setup(t, tt.err)
				err := s.Set(ctx)
				assert.ErrorIs(t, err, tt.err)
			},
		)
	}
}
