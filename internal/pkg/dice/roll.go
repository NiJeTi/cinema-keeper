package dice

import (
	"math/rand/v2"
)

func Roll(size Size) int {
	//nolint:gosec // no need to use secure random
	return rand.IntN(size.Int()) + 1
}
