package utils

import (
	"math/rand/v2"
)

func RandomColor() int {
	const maxColorValue = 0xffffff
	//nolint:gosec // no need to use secure random
	return rand.IntN(maxColorValue + 1)
}

func Ptr[T any](v T) *T {
	return &v
}
