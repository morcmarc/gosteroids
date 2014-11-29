package graphics

import (
	"github.com/go-gl/gl"
)

const bg_vertex = `#version 330

layout(location = 0) in vec3 vertexPosition_modelspace;
layout(location = 1) in vec2 vertexUV;
out vec2 UV;

void main()
{
    gl_Position = vec4(position, 1.0);
    UV = vertexUV;
}`

const bg_fragment = `#version 330

uniform sampler2D texSampler;
in vec2 UV;
out vec3 color;

void main() {
    color = texture2D(texSampler, UV).rgb;
}`

type Background struct {
	SceneObject
	Texture  gl.Texture
	Vertices []float32
	Vao      gl.VertexArray
	Vbo      gl.Buffer
	Ubo      gl.Buffer
	Sampler  gl.UniformLocation
	Program  gl.Program
}

func NewBackground() *Background {
	bg := &Background{
		Texture: OpenImageAsTexture("assets/nebula.png"),
	}

	bg.Vao = gl.GenVertexArray()
	bg.Vao.Bind()

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

	// ---------------------------------------------

	bg.Sampler = bg.Program.GetUniformLocation("texSampler")
	bg.Program.BindFragDataLocation(0, "color")

	uvBufferData := [12]float32{
		// Left bottom triangle
		0.0, 1.0,
		0.0, 0.0,
		1.0, 0.0,
		// Right top triangle
		1.0, 0.0,
		1.0, 1.0,
		0.0, 1.0,
	}

	bg.Ubo = gl.GenBuffer()
	bg.Ubo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvBufferData)*4, &uvBufferData, gl.STATIC_DRAW)

	vertices := [18]float32{
		// Left bottom triangle
		-1.0, 1.0, -1.0,
		-1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		// Right top triangle
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,
	}
	bg.Vbo = gl.GenBuffer()
	bg.Vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, &vertices, gl.STATIC_DRAW)

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

	b.Sampler.Uniform1i(0)

	vertexAttrib := gl.AttribLocation(0)
	vertexAttrib.EnableArray()
	defer vertexAttrib.DisableArray()
	b.Vbo.Bind(gl.ARRAY_BUFFER)
	defer b.Vbo.Unbind(gl.ARRAY_BUFFER)
	vertexAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

	uvAttrib := gl.AttribLocation(1)
	uvAttrib.EnableArray()
	defer uvAttrib.DisableArray()
	b.Ubo.Bind(gl.ARRAY_BUFFER)
	defer b.Ubo.Unbind(gl.ARRAY_BUFFER)
	uvAttrib.AttribPointer(2, gl.FLOAT, false, 0, nil)

	gl.DrawArrays(gl.TRIANGLES, 0, 3*2)
}
