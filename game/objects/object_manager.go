package objects

import (
	"math"

	b "github.com/morcmarc/gosteroids/game/broadcast"
	. "github.com/morcmarc/gosteroids/game/shared"
)

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

func (o *ObjectManager) Reset() {
	o.Spaceship.Reset()
	for _, a := range o.Asteroids {
		a.Reset()
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

func (o *ObjectManager) CheckCollision() bool {
	for _, a := range o.Asteroids {
		var dx float64 = (o.Spaceship.Position[0] + o.Spaceship.Radius) - (a.Position[0] + a.Radius)
		var dy float64 = (o.Spaceship.Position[1] + o.Spaceship.Radius) - (a.Position[1] + a.Radius)
		var distance float64 = math.Sqrt(dx*dx + dy*dy)

		if distance < o.Spaceship.Radius+a.Radius {
			return true
		}
	}
	return false
}

func (o *ObjectManager) CheckHits() (int, int) {
	for i, p := range o.Projectiles {
		for j, a := range o.Asteroids {
			var dx float64 = (p.Position[0] + 0.003) - (a.Position[0] + a.Radius)
			var dy float64 = (p.Position[1] + 0.003) - (a.Position[1] + a.Radius)
			var distance float64 = math.Sqrt(dx*dx + dy*dy)

			if distance < 0.003+a.Radius {
				return i, j
			}
		}
	}
	return -1, -1
}
