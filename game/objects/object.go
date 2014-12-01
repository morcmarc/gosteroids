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

func (o *ObjectManager) Listen(controlChanel chan uint8) {
	for m := range controlChanel {
		if m == Throttle {
			o.Spaceship.Position[1] += 0.01
		}
		if m == Break {
			o.Spaceship.Position[1] -= 0.01
		}
		if m == Left {
			o.Spaceship.Position[2] -= 0.1
		}
		if m == Right {
			o.Spaceship.Position[2] += 0.1
		}
	}
}
