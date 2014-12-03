package objects

import (
	"math"

	. "github.com/morcmarc/gosteroids/game/shared"
)

const (
	AccelerationCoeff float64 = 0.0010
	DecelerationCoeff float64 = 0.0010
	SlowdownCoeff     float64 = 0.9700
)

type Spaceship struct {
	Object
	Position [3]float64
	Velocity [2]float64
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Position: [3]float64{0.0, 0.0, 0.0},
		Velocity: [2]float64{0.0, 0.0},
	}
	return ss
}

func (s *Spaceship) Update() {
	s.slowDown()

	s.Position[0] += s.Velocity[0]
	s.Position[1] += s.Velocity[1]

	if s.Position[0] > 1.0 {
		s.Position[0] = -1.0
	}
	if s.Position[0] < -1.0 {
		s.Position[0] = 1.0
	}
	if s.Position[1] > 1.0 {
		s.Position[1] = -1.0
	}
	if s.Position[1] < -1.0 {
		s.Position[1] = 1.0
	}
}

func (s *Spaceship) Rotate(dir uint8) {
	switch dir {
	case Left:
		s.Position[2] -= 0.052359 // 3 degrees
		break
	case Right:
		s.Position[2] += 0.052359 // 3 degrees
		break
	}
}

func (s *Spaceship) Accelerate() {
	rx, ry := s.getRotationVector()
	s.Velocity[0] += AccelerationCoeff * rx
	s.Velocity[1] += AccelerationCoeff * ry
}

func (s *Spaceship) Decelerate() {
	rx, ry := s.getRotationVector()
	s.Velocity[0] -= DecelerationCoeff * rx
	s.Velocity[1] -= DecelerationCoeff * ry
}

func (s *Spaceship) slowDown() {
	s.Velocity[0] *= SlowdownCoeff
	s.Velocity[1] *= SlowdownCoeff
}

func (s *Spaceship) getRotationVector() (float64, float64) {
	return math.Sin(s.Position[2]), math.Cos(s.Position[2])
}
