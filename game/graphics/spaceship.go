package graphics

import (
	"github.com/go-gl/gl"
	o "github.com/morcmarc/gosteroids/game/objects"
)

const ss_vertex = `#version 330

in vec2 vrtx;
uniform vec3 position;

void main()
{
	float x = vrtx[0];
    float y = vrtx[1];
    float x_pos = position[0];
    float y_pos = position[1];
    float angle = position[2];
    float xx = (x * cos(angle) + y * sin(angle)) + x_pos;
    float yy = (-x * sin(angle) + y * cos(angle)) + y_pos;
    gl_Position = vec4(xx, yy, 0.0, 1.0);
}`

const ss_fragment = `#version 330

out vec4 outColor;

void main()
{
    outColor = vec4(1.0, 1.0, 1.0, 1.0);
}`

type Spaceship struct {
	SceneObject
	SSObject *o.Spaceship
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewSpaceship(sso *o.Spaceship) *Spaceship {
	ss := &Spaceship{
		Vertices: []float32{
			-0.02, -0.05,
			0.02, -0.05,
			0.0, 0.05,
		},
		SSObject: sso,
	}

	ss.Vbo = gl.GenBuffer()
	ss.Vbo.Bind(gl.ARRAY_BUFFER)

	ss.Vao = gl.GenVertexArray()
	ss.Vao.Bind()
	defer ss.Vao.Unbind()

	vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
	vertex_shader.Source(ss_vertex)
	vertex_shader.Compile()

	fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragment_shader.Source(ss_fragment)
	fragment_shader.Compile()

	ss.Program = gl.CreateProgram()
	ss.Program.AttachShader(vertex_shader)
	ss.Program.AttachShader(fragment_shader)

	ss.Program.Link()
	ss.Program.Use()
	defer ss.Program.Unuse()
	ss.Program.BindFragDataLocation(0, "outColor")

	vrtx := ss.Program.GetAttribLocation("vrtx")
	vrtx.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vrtx.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(ss.Vertices)*4, ss.Vertices, gl.DYNAMIC_DRAW)

	return ss
}

func (s *Spaceship) Draw() {
	s.Program.Use()
	defer s.Program.Unuse()

	s.Vao.Bind()
	defer s.Vao.Unbind()

	p := s.Program.GetUniformLocation("position")
	p.Uniform3f(
		float32(s.SSObject.Position[0]),
		float32(s.SSObject.Position[1]),
		float32(s.SSObject.Position[2]))

	gl.DrawArrays(gl.TRIANGLES, 0, len(s.Vertices))
}
