package objects

import (
	"math/rand"
)

type Asteroid struct {
	Object
	Position [3]float64
	Velocity [2]float64
}

func NewAsteroid() *Asteroid {
	x := rand.Float64()*(1-(-1)) + (-1)
	y := rand.Float64()*(1-(-1)) + (-1)
	vx := rand.Float64()*(1-(-1)) + (-1)
	vy := rand.Float64()*(1-(-1)) + (-1)
	a := &Asteroid{
		Position: [3]float64{x, y, 0.0},
		Velocity: [2]float64{vx / 500, vy / 500},
	}
	return a
}

func (a *Asteroid) Update() {
	a.Position[0] += a.Velocity[0]
	a.Position[1] += a.Velocity[1]

	if a.Position[0] > 1.0 {
		a.Position[0] = -1.0
	}
	if a.Position[0] < -1.0 {
		a.Position[0] = 1.0
	}
	if a.Position[1] > 1.0 {
		a.Position[1] = -1.0
	}
	if a.Position[1] < -1.0 {
		a.Position[1] = 1.0
	}
}
