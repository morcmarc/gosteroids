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
	Score         *Score
}

type SceneObject interface {
	Draw(ct float32)
	Delete()
}

func NewScene(om *o.ObjectManager, w, h, bgQuality int) *Scene {
	s := &Scene{
		ObjectManager: om,
		Spaceship:     NewSpaceship(om.Spaceship),
		Background:    NewBackground(bgQuality),
		Score:         NewScore(w, h),
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
	s.Score.Points += 1
	// TODO: remove indirect reference
	for i, p := range s.Projectiles {
		if p == nil {
			continue
		}
		if p.PSObject.IsOffScreen() {
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
	s.Score.Draw(ct)
}

func (s *Scene) Delete() {
	s.Background.Delete()
	s.Spaceship.Delete()
	for _, a := range s.Asteroids {
		a.Delete()
	}
	for _, p := range s.Projectiles {
		p.Delete()
	}
	s.Score.Delete()
}
