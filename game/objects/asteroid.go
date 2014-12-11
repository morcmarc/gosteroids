package objects

import (
	. "github.com/morcmarc/gosteroids/game/shared"
	"github.com/satori/go.uuid"
)

type Asteroid struct {
	Object
	Id       uuid.UUID
	Radius   float64
	Position [3]float64
	Velocity [3]float64
}

func NewAsteroid() *Asteroid {
	a := &Asteroid{}
	a.Reset()
	return a
}

func (a *Asteroid) Reset() {
	r := RandFloat64nInRange(0.0, 1.0)/20.0 + 0.05
	x := RandFloat64nInRange(-1.0, 1.0)
	y := RandFloat64nInRange(-1.0, 1.0)
	vx := RandFloat64nInRange(-1.0, 1.0)
	vy := RandFloat64nInRange(-1.0, 1.0)
	vr := RandFloat64nInRange(-1.0, 1.0)

	a.Id = uuid.NewV4()
	a.Radius = r
	a.Position = [3]float64{x, y, 0.0}
	a.Velocity = [3]float64{vx / 250, vy / 250, vr / 50}
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
