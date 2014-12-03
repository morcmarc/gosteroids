package graphics

import (
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Scene struct {
	ObjectManager *o.ObjectManager
	Background    *Background
	Spaceship     *Spaceship
	Asteroids     []*Asteroid
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
	}

	for _, ao := range om.Asteroids {
		a := NewAsteroid(ao)
		s.Asteroids = append(s.Asteroids, a)
	}

	return s
}

func (s *Scene) Update(ct float32) {
	s.ObjectManager.Update()
}

func (s *Scene) Draw(ct float32) {
	s.Background.Draw(ct)
	s.Spaceship.Draw(ct)

	for _, a := range s.Asteroids {
		a.Draw(ct)
	}
}
