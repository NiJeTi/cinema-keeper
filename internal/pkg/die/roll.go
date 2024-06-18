package die

import (
	"math/rand/v2"
)

func Roll(size Size) int {
	return rand.IntN(size.Int()) + 1
}
