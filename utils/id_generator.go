package utils

import "math/rand"

type IdGenerator func() string

func NewIdGenerator() string {
	return string(rand.Uint32())
}
