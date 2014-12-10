package shared

import (
	"math/rand"
	"time"
)

func RandFloat64nInRange(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-(min)) + (min)
}

func GetNormals(a, b [2]float64) ([2]float64, [2]float64) {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	n1 := [2]float64{-dy, dx}
	n2 := [2]float64{dy, -dx}
	return n1, n2
}
