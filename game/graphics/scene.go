package graphics

import (
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Scene struct {
	ObjectManager *o.ObjectManager
	Background    *Background
	Spaceship     *Spaceship
}

type SceneObject interface {
	Draw(ct float32)
}

func NewScene(om *o.ObjectManager) *Scene {
	s := &Scene{
		ObjectManager: om,
		Spaceship:     NewSpaceship(om.Spaceship),
		Background:    NewBackground(),
	}

	return s
}

func (s *Scene) Update(ct float32) {
	s.ObjectManager.Update()
}

func (s *Scene) Draw(ct float32) {
	s.Background.Draw(ct)
	s.Spaceship.Draw(ct)
}
