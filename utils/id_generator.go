package utils

import (
	"math/rand"
	"strconv"
)

type IdGenerator func() string

func NewIdGenerator() string {
	return strconv.Itoa(int(rand.Uint32()))
}
