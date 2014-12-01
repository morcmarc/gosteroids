package graphics

import (
	"github.com/go-gl/gl"
)

const bg_vertex = `#version 330

in vec2 vrtx;
out vec2 out_coord;

void main()
{
	float x = vrtx[0];
    float y = vrtx[1];
    // Translate: 
    // -1 -> 0
    //  0 -> 0.5
    //  1 -> 1
    out_coord = vec2(x / 2 + 0.5, y / 2 + 0.5);
    gl_Position = vec4(x, y, 0.9, 1.0);
}`

const bg_fragment = `#version 330

uniform sampler2D texSampler;
in vec2 out_coord;
out vec4 outColor;

void main()
{
    // outColor = vec4(texture2D(texSampler, out_coord).rgb, 1.0);
    outColor = vec4(out_coord[0], out_coord[1], 0.70, 1.0);
}`

type Background struct {
	SceneObject
	Texture  gl.Texture
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Program  gl.Program
}

func NewBackground() *Background {
	bg := &Background{
		Texture: OpenImageAsTexture("assets/nebula.png"),
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

	vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
	vertex_shader.Source(bg_vertex)
	vertex_shader.Compile()

	fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragment_shader.Source(bg_fragment)
	fragment_shader.Compile()

	bg.Program = gl.CreateProgram()
	bg.Program.AttachShader(vertex_shader)
	bg.Program.AttachShader(fragment_shader)

	bg.Program.Link()
	bg.Program.Use()
	// defer bg.Program.Unuse()
	bg.Program.BindFragDataLocation(0, "outColor")

	vrtx := bg.Program.GetAttribLocation("vrtx")
	vrtx.AttribPointer(2, gl.FLOAT, false, 0, nil)
	vrtx.EnableArray()

	gl.BufferData(gl.ARRAY_BUFFER, len(bg.Vertices)*4, bg.Vertices, gl.DYNAMIC_DRAW)

	return bg
}

func (b *Background) Draw() {
	b.Program.Use()
	defer b.Program.Unuse()

	b.Vao.Bind()
	defer b.Vao.Unbind()

	gl.ActiveTexture(gl.TEXTURE0)
	b.Texture.Bind(gl.TEXTURE_2D)
	defer b.Texture.Unbind(gl.TEXTURE_2D)

	s := b.Program.GetUniformLocation("texSampler")
	s.Uniform1i(0)

	gl.DrawArrays(gl.TRIANGLES, 0, len(b.Vertices))
}
