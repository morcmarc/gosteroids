package objects

import (
	"math"

	"github.com/satori/go.uuid"
)

const (
	ProjectileSpeedCoeff float64 = 0.03
)

type Projectile struct {
	Object
	Id       uuid.UUID
	Position [3]float64
	Velocity [2]float64
}

func NewProjectile(pos [3]float64) *Projectile {
	p := &Projectile{
		Id:       uuid.NewV4(),
		Position: pos,
		Velocity: [2]float64{0.0, 0.0},
	}

	rx, ry := p.getRotationVector()
	p.Velocity[0] += ProjectileSpeedCoeff * rx
	p.Velocity[1] += ProjectileSpeedCoeff * ry

	return p
}

func (p *Projectile) Update() {
	p.Position[0] += p.Velocity[0]
	p.Position[1] += p.Velocity[1]
}

func (p *Projectile) IsOffScreen() bool {
	if p.Position[0] > 1.0 || p.Position[0] < -1.0 || p.Position[1] > 1.0 || p.Position[1] < -1.0 {
		return true
	}
	return false
}

func (p *Projectile) getRotationVector() (float64, float64) {
	return math.Sin(p.Position[2]), math.Cos(p.Position[2])
}
