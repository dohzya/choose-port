package chooseport

import (
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

var DEFAULT_FROM = 2000
var DEFAULT_TO = 40000

func Random() int {
	return Between(DEFAULT_FROM, DEFAULT_TO)
}

func From(from int) int {
	return Between(from, from+1000)
}

func Between(from int, to int) int {
	return seed.Intn(to-from) + from
}
