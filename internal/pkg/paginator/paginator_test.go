package paginator_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nijeti/cinema-keeper/internal/pkg/paginator"
)

func TestInfo(t *testing.T) {
	t.Parallel()

	tests := map[string]func(t *testing.T){
		"first_page": func(t *testing.T) {
			info := paginator.Info(100, 10, 0)
			assert.Equal(
				t, paginator.PageInfo{
					Page:     0,
					LastPage: 9,
					Offset:   0,
					Limit:    9,
				},
				info,
			)
		},
		"last_page": func(t *testing.T) {
			info := paginator.Info(100, 10, 90)
			assert.Equal(
				t, paginator.PageInfo{
					Page:     9,
					LastPage: 9,
					Offset:   90,
					Limit:    99,
				},
				info,
			)
		},
		"uneven_elems": func(t *testing.T) {
			info := paginator.Info(11, 10, 10)
			assert.Equal(
				t, paginator.PageInfo{
					Page:     1,
					LastPage: 1,
					Offset:   10,
					Limit:    10,
				},
				info,
			)
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func TestPaginate(t *testing.T) {
	t.Parallel()

	tests := map[string]func(t *testing.T){
		"success": func(t *testing.T) {
			pages := make([][]int, 0, 3)
			infos := make([]paginator.PageInfo, 0, 3)

			err := paginator.Paginate(
				[]int{1, 2, 3, 4, 5},
				2,
				func(page []int, info paginator.PageInfo) error {
					pages = append(pages, page)
					infos = append(infos, info)
					return nil
				},
			)

			assert.NoError(t, err)

			assert.Len(t, pages, 3)
			assert.Len(t, infos, 3)

			assert.Equal(t, []int{1, 2}, pages[0])
			assert.Equal(
				t, paginator.PageInfo{
					Page: 0, LastPage: 2, Offset: 0, Limit: 1,
				},
				infos[0],
			)

			assert.Equal(t, []int{3, 4}, pages[1])
			assert.Equal(
				t, paginator.PageInfo{
					Page: 1, LastPage: 2, Offset: 2, Limit: 3,
				},
				infos[1],
			)

			assert.Equal(t, []int{5}, pages[2])
			assert.Equal(
				t, paginator.PageInfo{
					Page: 2, LastPage: 2, Offset: 4, Limit: 4,
				},
				infos[2],
			)
		},
		"error": func(t *testing.T) {
			wantErr := errors.New("error")
			pageCount := 0

			err := paginator.Paginate(
				[]int{1, 2, 3, 4, 5}, 2,
				func(_ []int, _ paginator.PageInfo) error {
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
