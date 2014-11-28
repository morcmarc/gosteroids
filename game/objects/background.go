package objects

import (
	"github.com/go-gl/gl"
	"github.com/morcmarc/gosteroids/game/utils"
)

type Background struct {
	Object
	Texture gl.Texture
}

func NewBackground() *Background {
	bg := &Background{
		Texture: utils.OpenImageAsTexture("assets/nebula.png"),
	}
	return bg
}

func (b *Background) Draw() {
	gl.Enable(gl.TEXTURE_2D)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)

	vertices := []float32{
		-1.0, 1.0, 0.0,
		1.0, 1.0, 0.0,
		1.0, -1.0, 0.0,
		-1.0, -1.0, 0.0,
	}
	texVertices := []float32{
		0.0, 0.0, 0.0,
		1.0, 0.0, 0.0,
		1.0, 1.0, 0.0,
		0.0, 1.0, 0.0,
	}

	gl.VertexPointer(3, gl.FLOAT, 0, &vertices)
	gl.TexCoordPointer(3, gl.FLOAT, 0, &texVertices)

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, vertices, gl.STATIC_DRAW)
	gl.DrawArrays(gl.QUADS, 0, 4)

	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.Disable(gl.TEXTURE_2D)
}
