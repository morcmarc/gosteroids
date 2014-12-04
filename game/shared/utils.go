package shared

import (
	"math/rand"
	"time"
)

func RandFloat64nInRange(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-(min)) + (min)
}
