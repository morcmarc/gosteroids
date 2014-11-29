package objects

import ()

type Spaceship struct {
	Object
	Position []float32
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Position: []float32{0.0, 0.0, 0.0},
	}
	return ss
}
