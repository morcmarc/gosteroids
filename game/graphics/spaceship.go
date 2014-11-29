package graphics

import (
	"github.com/go-gl/gl"
)

const ss_vertex = `#version 330

in vec2 position;

void main()
{
    gl_Position = vec4(position, 0.0, 1.0);
}`

const ss_fragment = `#version 330

out vec4 outColor;

void main()
{
    outColor = vec4(1.0, 1.0, 1.0, 1.0);
}`

type Spaceship struct {
	SceneObject
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewSpaceship() *Spaceship {
	ss := &Spaceship{
		Vertices: []float32{
			-0.05, 0.0, 0.0,
			0.05, 0.0, 0.0,
			0.0, 0.1, 0.0,
		},
	}

	ss.Vbo = gl.GenBuffer()
	ss.Vbo.Bind(gl.ARRAY_BUFFER)

	ss.Vao = gl.GenVertexArray()
	ss.Vao.Bind()

	gl.BufferData(gl.ARRAY_BUFFER, len(ss.Vertices)*4, ss.Vertices, gl.STATIC_DRAW)

	vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
	vertex_shader.Source(ss_vertex)
	vertex_shader.Compile()

	fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragment_shader.Source(ss_fragment)
	fragment_shader.Compile()

	ss.Program = gl.CreateProgram()
	ss.Program.AttachShader(vertex_shader)
	ss.Program.AttachShader(fragment_shader)

	ss.Program.BindFragDataLocation(0, "outColor")
	ss.Program.Link()
	ss.Program.Use()

	positionAttrib := ss.Program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()

	ss.Vao.Unbind()

	return ss
}

func (s *Spaceship) Draw() {
	gl.EnableClientState(gl.VERTEX_ARRAY)
	s.Program.Use()
	s.Vao.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	s.Vao.Unbind()
	s.Program.Unuse()
	gl.DisableClientState(gl.VERTEX_ARRAY)
}