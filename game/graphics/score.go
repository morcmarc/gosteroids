package graphics

import (
	"fmt"
)

type Score struct {
	SceneObject
	Points int
	Font   *Font
}

func NewScore() *Score {
	s := &Score{
		Points: 0,
	}

	s.Font = NewFont("assets/fonts/alphabet_30.png", 16, 6, 600, 600)

	return s
}

func (s *Score) Draw(ct float32) {
	ps := fmt.Sprintf("%06d", s.Points)
	s.Font.Printf(ps, -0.9, -0.8)
}
