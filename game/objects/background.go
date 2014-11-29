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
	b.Texture.Bind(gl.TEXTURE_2D)

	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, -1, 1)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, -1, 1)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, 1, 1)
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, 1, 1)

	gl.End()
}
