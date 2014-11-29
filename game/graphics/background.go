package graphics

import (
	"github.com/go-gl/gl"
)

const bg_vertex = `#version 330

in vec2 position;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main()
{
	fragTexCoord = vertTexCoord;
    gl_Position = vec4(position, 1.0, 1.0);
}`

const bg_fragment = `#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
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
		Texture: OpenImageAsTexture("assets/nebula.png", gl.TEXTURE0),
		Vertices: []float32{
			// Left bottom triangle
			-1.0, 1.0, 0.0, 0.0, 0.0,
			-1.0, -1.0, 0.0, 1.0, 0.0,
			1.0, -1.0, 0.0, 0.0, 1.0,
			// Right top triangle
			1.0, -1.0, 0.0, 1.0, 1.0,
			1.0, 1.0, 0.0, 1.0, 0.0,
			-1.0, 1.0, 0.0, 0.0, 0.0,
		},
	}

	bg.Vbo = gl.GenBuffer()
	bg.Vbo.Bind(gl.ARRAY_BUFFER)

	bg.Vao = gl.GenVertexArray()
	bg.Vao.Bind()

	gl.BufferData(gl.ARRAY_BUFFER, len(bg.Vertices)*4, bg.Vertices, gl.STATIC_DRAW)

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

	texUni := bg.Program.GetUniformLocation("tex")
	texUni.Uniform1i(0)

	bg.Program.BindFragDataLocation(0, "outputColor")

	vertAttr := bg.Program.GetAttribLocation("vert")
	vertAttr.EnableArray()
	gl.VertexPointer(3, gl.FLOAT, 0, bg.Vertices)

	texAttr := bg.Program.GetAttribLocation("vertTexCoord")
	texAttr.EnableArray()
	gl.TexCoordPointer(2, gl.FLOAT, 0, bg.Vertices)

	bg.Vao.Unbind()

	return bg
}

func (b *Background) Draw() {
	gl.EnableClientState(gl.VERTEX_ARRAY)
	b.Program.Use()
	b.Vao.Bind()

	gl.ActiveTexture(gl.TEXTURE0)
	b.Texture.Bind(gl.TEXTURE_2D)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	b.Texture.Unbind(gl.TEXTURE_2D)

	b.Vao.Unbind()
	b.Program.Unuse()
	gl.DisableClientState(gl.VERTEX_ARRAY)
}
