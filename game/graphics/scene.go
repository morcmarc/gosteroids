package graphics

import (
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Scene struct {
	Background *Background
	Spaceship  *Spaceship
}

type SceneObject interface {
	Draw()
}

func NewScene(om *o.ObjectManager) *Scene {
	s := &Scene{
		Spaceship:  NewSpaceship(om.Spaceship),
		Background: NewBackground(),
	}
	return s
}

func (s *Scene) Draw() {
	s.Background.Draw()
	s.Spaceship.Draw()
}
