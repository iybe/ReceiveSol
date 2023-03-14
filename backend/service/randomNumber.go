package service

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(90000) + 10000
	return number
}
