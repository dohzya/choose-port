package chooseport

import (
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func ChoosePort(from int, to int) int {
	return seed.Intn(to-from) + from
}
