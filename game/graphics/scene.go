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
		Background: NewBackground(),
		Spaceship:  NewSpaceship(),
	}
	return s
}

func (s *Scene) Draw() {
	s.Background.Draw()
	s.Spaceship.Draw()
}
