package graphics

import (
// "github.com/morcmarc/gosteroids/game/objects"
)

type Scene struct {
	Background *Background
	Spaceship  *Spaceship
}

type SceneObject interface {
	Draw()
}

func NewScene() *Scene {
	s := &Scene{
		Spaceship:  NewSpaceship(),
		Background: NewBackground(),
	}
	return s
}

func (s *Scene) Draw() {
	s.Background.Draw()
	s.Spaceship.Draw()
}
