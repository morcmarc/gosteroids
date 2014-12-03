package graphics

import (
	"github.com/go-gl/gl"
)

type Background struct {
	SceneObject
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewBackground() *Background {
	bg := &Background{
		Vertices: []float32{
			// Left bottom triangle
			-1.0, 1.0,
			-1.0, -1.0,
			1.0, -1.0,
			// Right top triangle
			1.0, -1.0,
			1.0, 1.0,
			-1.0, 1.0,
		},
	}

	bg.Vbo = gl.GenBuffer()
	bg.Vbo.Bind(gl.ARRAY_BUFFER)

	bg.Vao = gl.GenVertexArray()
	bg.Vao.Bind()
	defer bg.Vao.Unbind()

	vertexShader, err := LoadShader("assets/shaders/background.vertex.glsl", VertexShader)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := LoadShader("assets/shaders/background.fragment.glsl", FragmentShader)
	if err != nil {
		panic(err)
	}

	bg.Program = NewProgram(vertexShader, fragmentShader)
	bg.Program.Use()
	defer bg.Program.Unuse()
	bg.Program.BindFragDataLocation(0, "outColor")

	vrtx := bg.Program.GetAttribLocation("vrtx")
	vrtx.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vrtx.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(bg.Vertices)*4, bg.Vertices, gl.DYNAMIC_DRAW)

	return bg
}

func (b *Background) Draw(ct float32) {
	b.Program.Use()
	defer b.Program.Unuse()

	b.Vao.Bind()
	defer b.Vao.Unbind()

	t := b.Program.GetUniformLocation("time")
	t.Uniform1f(ct)

	r := b.Program.GetUniformLocation("resolution")
	r.Uniform2f(600.0, 600.0)

	gl.DrawArrays(gl.TRIANGLES, 0, len(b.Vertices))
}
