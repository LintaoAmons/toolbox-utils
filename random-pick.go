package toolboxutils

import (
	"math/rand"
	"time"
)

func RandomString(options []string) string {
	rand.Seed(time.Now().UnixNano())
	return options[rand.Intn(len(options))]
}
