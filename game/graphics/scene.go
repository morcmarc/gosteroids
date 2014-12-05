package graphics

import (
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Scene struct {
	ObjectManager *o.ObjectManager
	Background    *Background
	Spaceship     *Spaceship
	Asteroids     []*Asteroid
	Projectiles   []*Projectile
}

type SceneObject interface {
	Draw(ct float32)
}

func NewScene(om *o.ObjectManager) *Scene {
	s := &Scene{
		ObjectManager: om,
		Spaceship:     NewSpaceship(om.Spaceship),
		Background:    NewBackground(),
		Asteroids:     []*Asteroid{},
		Projectiles:   []*Projectile{},
	}

	for _, ao := range om.Asteroids {
		a := NewAsteroid(ao)
		s.Asteroids = append(s.Asteroids, a)
	}

	return s
}

func (s *Scene) Fire() {
	po := s.ObjectManager.FireProjectile()
	p := NewProjectile(po)
	// TODO: remove indirect reference
	s.ObjectManager.Projectiles = append(s.ObjectManager.Projectiles, po)
	s.Projectiles = append(s.Projectiles, p)
}

func (s *Scene) Update(ct float32) {
	s.ObjectManager.Update()
	// TODO: remove indirect reference
	for i, p := range s.Projectiles {
		if p == nil {
			continue
		}
		if p.SSObject.IsOffScreen() {
			copy(s.Projectiles[i:], s.Projectiles[i+1:])
			s.Projectiles[len(s.Projectiles)-1] = nil
			s.Projectiles = s.Projectiles[:len(s.Projectiles)-1]
		}
	}
}

func (s *Scene) Draw(ct float32) {
	s.Background.Draw(ct)
	s.Spaceship.Draw(ct)
	for _, a := range s.Asteroids {
		a.Draw(ct)
	}
	for _, p := range s.Projectiles {
		p.Draw(ct)
	}
}
