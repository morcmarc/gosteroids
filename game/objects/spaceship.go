package objects

import ()

type Spaceship struct {
	Object
	Position [3]float32
	Rotation [3]float32
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Position: [3]float32{0.0, 0.0, 0.0},
		Rotation: [3]float32{0.0, 0.0, 0.0},
	}
	return ss
}
