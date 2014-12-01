package objects

import (
	. "github.com/morcmarc/gosteroids/game/shared"
)

type Object interface {
	Update()
}

type ObjectManager struct {
	Spaceship *Spaceship
}

func NewObjectManager() *ObjectManager {
	om := &ObjectManager{
		Spaceship: NewSpaceship(),
	}
	return om
}

func (o *ObjectManager) Update() {
	o.Spaceship.Update()
}

func (o *ObjectManager) Listen(controlChanel chan uint8) {
	for m := range controlChanel {
		if m == Throttle {
			o.Spaceship.Accelerate()
		}
		if m == Break {
			o.Spaceship.Decelerate()
		}
		if m == Left {
			o.Spaceship.Rotate(Left)
		}
		if m == Right {
			o.Spaceship.Rotate(Right)
		}
	}
}
