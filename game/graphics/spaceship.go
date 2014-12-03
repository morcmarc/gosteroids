package graphics

import (
	"github.com/go-gl/gl"
	o "github.com/morcmarc/gosteroids/game/objects"
)

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
			0.0, 0.05,
			0.025, -0.05,
			0.0, -0.025,

			0.0, -0.025,
			-0.025, -0.05,
			0.0, 0.05,
		},
		SSObject: sso,
	}

	ss.Vbo = gl.GenBuffer()
	ss.Vbo.Bind(gl.ARRAY_BUFFER)

	ss.Vao = gl.GenVertexArray()
	ss.Vao.Bind()
	defer ss.Vao.Unbind()

	vertexShader, err := LoadShader("assets/shaders/spaceship.vertex.glsl", VertexShader)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := LoadShader("assets/shaders/spaceship.fragment.glsl", FragmentShader)
	if err != nil {
		panic(err)
	}

	ss.Program = NewProgram(vertexShader, fragmentShader)
	ss.Program.Use()
	defer ss.Program.Unuse()
	ss.Program.BindFragDataLocation(0, "outColor")

	vrtx := ss.Program.GetAttribLocation("vrtx")
	vrtx.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vrtx.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(ss.Vertices)*4, ss.Vertices, gl.DYNAMIC_DRAW)

	return ss
}

func (s *Spaceship) Draw(ct float32) {
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
