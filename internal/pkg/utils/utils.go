package utils

import (
	"math/rand/v2"
)

func RandomColor() int {
	const maxColorValue = 0xffffff
	return rand.IntN(maxColorValue + 1)
}
