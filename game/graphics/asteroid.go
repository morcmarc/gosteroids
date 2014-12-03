package graphics

import (
	"math/rand"

	"github.com/go-gl/gl"
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Asteroid struct {
	SceneObject
	AObject  *o.Asteroid
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewAsteroid(ao *o.Asteroid) *Asteroid {
	max := float32(rand.Intn(200)) / 2000.0
	min := max / 2.0

	ss := &Asteroid{
		Vertices: []float32{
			-min, max,
			min, max,
			max, 0.0,

			max, 0.00,
			min, -max,
			-min, -max,

			-min, -max,
			-max, 0.00,
			-min, max,

			-min, max,
			max, 0.00,
			-min, -max,
		},
		AObject: ao,
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

func (s *Asteroid) Draw(ct float32) {
	s.Program.Use()
	defer s.Program.Unuse()

	s.Vao.Bind()
	defer s.Vao.Unbind()

	p := s.Program.GetUniformLocation("position")
	p.Uniform3f(
		float32(s.AObject.Position[0]),
		float32(s.AObject.Position[1]),
		float32(s.AObject.Position[2]))

	gl.DrawArrays(gl.TRIANGLES, 0, len(s.Vertices))
}
