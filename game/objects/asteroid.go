package objects

import (
	. "github.com/morcmarc/gosteroids/game/shared"
)

type Asteroid struct {
	Object
	Position [3]float64
	Velocity [3]float64
}

func NewAsteroid() *Asteroid {
	x := RandFloat64nInRange(-1.0, 1.0)
	y := RandFloat64nInRange(-1.0, 1.0)
	vx := RandFloat64nInRange(-1.0, 1.0)
	vy := RandFloat64nInRange(-1.0, 1.0)
	vr := RandFloat64nInRange(-1.0, 1.0)

	a := &Asteroid{
		Position: [3]float64{x, y, 0.0},
		Velocity: [3]float64{vx / 500, vy / 500, vr / 100},
	}

	return a
}

func (a *Asteroid) Update() {
	a.Position[0] += a.Velocity[0]
	a.Position[1] += a.Velocity[1]
	a.Position[2] += a.Velocity[2]

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
