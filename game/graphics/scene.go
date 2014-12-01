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
	Draw()
}

func NewScene(om *o.ObjectManager) *Scene {
	s := &Scene{
		ObjectManager: om,
		Spaceship:     NewSpaceship(om.Spaceship),
		Background:    NewBackground(),
	}
	return s
}

func (s *Scene) Update() {
	s.ObjectManager.Update()
}

func (s *Scene) Draw() {
	s.Background.Draw()
	s.Spaceship.Draw()
}
