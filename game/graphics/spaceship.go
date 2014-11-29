package graphics

import (
	"github.com/go-gl/gl"
	// "github.com/morcmarc/gosteroids/game/objects"
)

type Spaceship struct {
	SceneObject
	Vertices []float32
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Vertices: []float32{
			-0.05, 0.0, 0.0,
			0.05, 0.0, 0.0,
			0.0, 0.1, 0.0,
		},
	}
	return ss
}

func (s *Spaceship) Draw() {
	gl.EnableClientState(gl.VERTEX_ARRAY)

	gl.VertexPointer(3, gl.FLOAT, 0, s.Vertices)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	gl.DisableClientState(gl.VERTEX_ARRAY)
}
