package graphics

import (
	"math"
	"math/rand"
	"sort"
	"time"

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
	ss := &Asteroid{
		Vertices: generateAsteroidVertices(25),
		AObject:  ao,
	}

	ss.Vbo = gl.GenBuffer()
	ss.Vbo.Bind(gl.ARRAY_BUFFER)

	ss.Vao = gl.GenVertexArray()
	ss.Vao.Bind()
	defer ss.Vao.Unbind()

	vertexShader, err := LoadShader("assets/shaders/generic.vertex.glsl", VertexShader)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := LoadShader("assets/shaders/generic.fragment.glsl", FragmentShader)
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

func generateAsteroidVertices(resolution int) []float32 {
	vertices := []float32{}

	rand.Seed(time.Now().UnixNano())

	radius := rand.Float64()/20.0 + 0.05
	angles := []int{}

	// Generate N points on a circle
	for i := 0; i < resolution; i++ {
		angles = append(angles, rand.Intn(360))
	}
	sort.Ints(angles)

	for i, a := range angles {
		var prev int
		if i == 0 {
			prev = angles[len(angles)-1]
		} else {
			prev = angles[i-1]
		}

		prevRadian := (float64(prev) * math.Pi) / 180.0
		prevX := 0.0 + radius*math.Cos(prevRadian)
		prevY := 0.0 + radius*math.Sin(prevRadian)

		radian := (float64(a) * math.Pi) / 180.0
		x := 0.0 + radius*math.Cos(radian)
		y := 0.0 + radius*math.Sin(radian)

		vertices = append(vertices, float32(0.0))
		vertices = append(vertices, float32(0.0))
		vertices = append(vertices, float32(prevX))
		vertices = append(vertices, float32(prevY))
		vertices = append(vertices, float32(x))
		vertices = append(vertices, float32(y))
	}

	return vertices
}
