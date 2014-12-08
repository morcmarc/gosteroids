package graphics

import (
	"github.com/go-gl/gl"
	o "github.com/morcmarc/gosteroids/game/objects"
)

type Projectile struct {
	SceneObject
	PSObject *o.Projectile
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewProjectile(pso *o.Projectile) *Projectile {
	ss := &Projectile{
		Vertices: []float32{
			-0.006, 0.00,
			0.000, 0.06,
			0.006, 0.00,
		},
		PSObject: pso,
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

func (p *Projectile) Draw(ct float32) {
	p.Program.Use()
	defer p.Program.Unuse()

	p.Vao.Bind()
	defer p.Vao.Unbind()

	pos := p.Program.GetUniformLocation("position")
	pos.Uniform3f(
		float32(p.PSObject.Position[0]),
		float32(p.PSObject.Position[1]),
		float32(p.PSObject.Position[2]))

	gl.DrawArrays(gl.TRIANGLES, 0, len(p.Vertices))
}

func (p *Projectile) Delete() {
	p.Vao.Delete()
	p.Vbo.Delete()
	p.Program.Delete()
}
