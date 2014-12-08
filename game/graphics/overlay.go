package graphics

import (
	"github.com/go-gl/gl"
)

type Overlay struct {
	SceneObject
	Width    int
	Height   int
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewOverlay(width, height int) *Overlay {
	bg := &Overlay{
		Width:  width,
		Height: height,
		Vertices: []float32{
			// Left bottom triangle
			-1.0, 1.0, 0.0, 0.0,
			-1.0, -1.0, 0.0, 0.0,
			1.0, -1.0, 0.0, 0.0,
			// Right top triangle
			1.0, -1.0, 0.0, 0.0,
			1.0, 1.0, 0.0, 0.0,
			-1.0, 1.0, 0.0, 0.0,
		},
	}

	vertexShader, err := LoadShader("assets/shaders/overlay.vertex.glsl", VertexShader)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := LoadShader("assets/shaders/overlay.fragment.glsl", FragmentShader)
	if err != nil {
		panic(err)
	}

	bg.Program = NewProgram(vertexShader, fragmentShader)
	bg.Program.Use()
	defer bg.Program.Unuse()

	bg.Vbo = gl.GenBuffer()
	bg.Vbo.Bind(gl.ARRAY_BUFFER)

	bg.Vao = gl.GenVertexArray()
	bg.Vao.Bind()
	defer bg.Vao.Unbind()

	gl.BufferData(gl.ARRAY_BUFFER, len(bg.Vertices)*4, bg.Vertices, gl.DYNAMIC_DRAW)

	return bg
}

func (b *Overlay) Draw(ct float32) {
	b.Program.Use()
	defer b.Program.Unuse()

	b.Vao.Bind()
	defer b.Vao.Unbind()

	b.Vbo.Bind(gl.ARRAY_BUFFER)
	defer b.Vbo.Unbind(gl.ARRAY_BUFFER)

	positionAttrib := b.Program.GetAttribLocation("position")
	positionAttrib.AttribPointer(4, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(b.Vertices)*4, b.Vertices, gl.DYNAMIC_DRAW)
	gl.DrawArrays(gl.TRIANGLES, 0, len(b.Vertices))
}

func (b *Overlay) Delete() {
	b.Vao.Delete()
	b.Vao.Delete()
	b.Program.Delete()
}
