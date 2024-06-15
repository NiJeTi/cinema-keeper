package utils

import (
	"math/rand/v2"
)

func RandomColor() int {
	const maxColorValue = 16777215
	return rand.IntN(maxColorValue + 1)
}
