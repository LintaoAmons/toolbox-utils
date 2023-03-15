package random

import (
	"math/rand"
	"time"
)

func RandomPick(options []string) string {
	rand.Seed(time.Now().UnixNano())
	return options[rand.Intn(len(options))]
}
