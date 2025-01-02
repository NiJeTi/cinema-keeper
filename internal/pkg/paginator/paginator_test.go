package paginator_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/pkg/paginator"
)

func TestPaginate(t *testing.T) {
	t.Parallel()

	tests := map[string]func(t *testing.T){
		"success": func(t *testing.T) {
			pages := make([][]int, 0, 3)

			err := paginator.Paginate(
				[]int{1, 2, 3, 4, 5}, 2, func(page []int) error {
					pages = append(pages, page)
					return nil
				},
			)

			assert.NoError(t, err)
			assert.Len(t, pages, 3)
			assert.Equal(t, []int{1, 2}, pages[0])
			assert.Equal(t, []int{3, 4}, pages[1])
			assert.Equal(t, []int{5}, pages[2])
		},
		"error": func(t *testing.T) {
			wantErr := errors.New("error")
			pageCount := 0

			err := paginator.Paginate(
				[]int{1, 2, 3, 4, 5}, 2, func(_ []int) error {
					pageCount++
					return wantErr
				},
			)

			assert.ErrorIs(t, err, wantErr)
			assert.Equal(t, 1, pageCount)
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}
