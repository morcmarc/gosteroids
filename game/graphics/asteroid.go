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
	a := &Asteroid{
		Vertices: generateAsteroidVertices(25),
		AObject:  ao,
	}

	vertexShader, err := LoadShader("assets/shaders/generic.vertex.glsl", VertexShader)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := LoadShader("assets/shaders/generic.fragment.glsl", FragmentShader)
	if err != nil {
		panic(err)
	}

	a.Program = NewProgram(vertexShader, fragmentShader)
	a.Program.Use()
	defer a.Program.Unuse()
	a.Program.BindFragDataLocation(0, "outColor")

	a.Vao = gl.GenVertexArray()
	a.Vao.Bind()

	a.Vbo = gl.GenBuffer()
	a.Vbo.Bind(gl.ARRAY_BUFFER)

	vrtx := a.Program.GetAttribLocation("vrtx")
	vrtx.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vrtx.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(a.Vertices)*4, a.Vertices, gl.DYNAMIC_DRAW)

	a.Vbo.Unbind(gl.ARRAY_BUFFER)
	a.Vao.Unbind()

	return a
}

func (a *Asteroid) Draw(ct float32) {
	a.Program.Use()
	defer a.Program.Unuse()

	a.Vao.Bind()
	defer a.Vao.Unbind()

	p := a.Program.GetUniformLocation("position")
	p.Uniform3f(
		float32(a.AObject.Position[0]),
		float32(a.AObject.Position[1]),
		float32(a.AObject.Position[2]))

	gl.DrawArrays(gl.TRIANGLES, 0, len(a.Vertices))
}

func (a *Asteroid) Delete() {
	a.Vao.Delete()
	a.Vbo.Delete()
	a.Program.Delete()
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

		vertices = append(vertices,
			float32(0.0),
			float32(0.0),
			float32(prevX),
			float32(prevY),
			float32(x),
			float32(y))
	}

	return vertices
}
