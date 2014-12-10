package objects

import (
	"math"

	b "github.com/morcmarc/gosteroids/game/broadcast"
	. "github.com/morcmarc/gosteroids/game/shared"
	"github.com/satori/go.uuid"
)

type ObjectManager struct {
	Spaceship   *Spaceship
	Asteroids   map[uuid.UUID]*Asteroid
	Projectiles map[uuid.UUID]*Projectile
}

func NewObjectManager() *ObjectManager {
	om := &ObjectManager{
		Spaceship:   NewSpaceship(),
		Asteroids:   map[uuid.UUID]*Asteroid{},
		Projectiles: map[uuid.UUID]*Projectile{},
	}

	om.Reset()

	return om
}

func (o *ObjectManager) Update() {
	o.Spaceship.Update()
	for _, a := range o.Asteroids {
		a.Update()
	}
	for _, p := range o.Projectiles {
		p.Update()
	}
}

func (o *ObjectManager) Reset() {
	o.Spaceship.Reset()

	for _, a := range o.Asteroids {
		o.RemoveAsteroid(a.Id)
	}

	for i := 0; i < 10; i++ {
		a := NewAsteroid()
		o.AddAsteroid(a)
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

func (o *ObjectManager) AddAsteroid(a *Asteroid) {
	o.Asteroids[a.Id] = a
}

func (o *ObjectManager) AddProjectile(p *Projectile) {
	o.Projectiles[p.Id] = p
}

func (o *ObjectManager) RemoveAsteroid(id uuid.UUID) {
	delete(o.Asteroids, id)
}

func (o *ObjectManager) RemoveProjectile(id uuid.UUID) {
	delete(o.Projectiles, id)
}

func (o *ObjectManager) FireProjectile() *Projectile {
	p := NewProjectile(o.Spaceship.Position)
	o.AddProjectile(p)
	return p
}

func (o *ObjectManager) CheckCollision() bool {
	for _, a := range o.Asteroids {
		// TODO: replace with Seperating Axis Theorem
		var dx float64 = (o.Spaceship.Position[0] + o.Spaceship.Radius) - (a.Position[0] + a.Radius)
		var dy float64 = (o.Spaceship.Position[1] + o.Spaceship.Radius) - (a.Position[1] + a.Radius)
		var distance float64 = math.Sqrt(dx*dx + dy*dy)

		if distance < o.Spaceship.Radius+a.Radius {
			return true
		}
	}
	return false
}

func (o *ObjectManager) CheckHits() (uuid.UUID, uuid.UUID) {
	for _, p := range o.Projectiles {
		for _, a := range o.Asteroids {
			// TODO: replace with Seperating Axis Theorem
			var dx float64 = (p.Position[0] + 0.003) - (a.Position[0] + a.Radius)
			var dy float64 = (p.Position[1] + 0.003) - (a.Position[1] + a.Radius)
			var distance float64 = math.Sqrt(dx*dx + dy*dy)

			if distance < 0.002+a.Radius {
				return p.Id, a.Id
			}
		}
	}
	return [16]byte{}, [16]byte{}
}
