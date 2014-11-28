package objects

import (
	"github.com/go-gl/gl"
)

type Spaceship struct {
	Object
	Vertices []float32
	Position []float32
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Vertices: []float32{
			-0.05, 0.0, 0.0,
			0.05, 0.0, 0.0,
			0.0, 0.1, 0.0,
		},
		Position: []float32{0.0, 0.0, 0.0},
	}
	return ss
}

func (s *Spaceship) Draw() {
	gl.BufferData(gl.ARRAY_BUFFER, len(s.Vertices)*4, s.Vertices, gl.STATIC_DRAW)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
