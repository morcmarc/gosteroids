package objects

import (
	b "github.com/morcmarc/gosteroids/game/broadcast"
	. "github.com/morcmarc/gosteroids/game/shared"
)

type Object interface {
	Update()
}

type ObjectManager struct {
	Spaceship   *Spaceship
	Asteroids   []*Asteroid
	Projectiles []*Projectile
}

func NewObjectManager() *ObjectManager {
	om := &ObjectManager{
		Spaceship:   NewSpaceship(),
		Asteroids:   []*Asteroid{},
		Projectiles: []*Projectile{},
	}

	for i := 0; i < 10; i++ {
		a := NewAsteroid()
		om.Asteroids = append(om.Asteroids, a)
	}

	return om
}

func (o *ObjectManager) Update() {
	o.Spaceship.Update()
	for _, a := range o.Asteroids {
		a.Update()
	}

	for i, p := range o.Projectiles {
		if p == nil {
			continue
		}
		if p.IsOffScreen() {
			copy(o.Projectiles[i:], o.Projectiles[i+1:])
			o.Projectiles[len(o.Projectiles)-1] = nil
			o.Projectiles = o.Projectiles[:len(o.Projectiles)-1]
		}
		if p != nil {
			p.Update()
		}
	}
}

func (o *ObjectManager) Listen(cc b.Receiver) {
	for m := cc.Read(); m != nil; m = cc.Read() {
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

func (o *ObjectManager) FireProjectile() *Projectile {
	p := NewProjectile(o.Spaceship.Position)
	return p
}
