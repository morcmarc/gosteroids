package objects

import (
	"fmt"
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
		fmt.Printf("Msg: %d\n", m)
	}
}
