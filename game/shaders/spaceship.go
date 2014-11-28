package shaders

import (
	"github.com/go-gl/gl"
)

const SpaceshipVertex string = `#version 330
in vec2 position;

void main()
{
    gl_Position = vec4(position, 0.0, 1.0);
}
`
const SpaceshipFragment string = `#version 330
out vec4 outColor;

void main()
{
    outColor = vec4(1.0, 1.0, 1.0, 1.0);
}
`

func GetSpaceshipVertex() gl.Shader {
	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vs.Source(SpaceshipVertex)
	vs.Compile()

	return vs
}

func GetSpaceshipFragment() gl.Shader {
	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fs.Source(SpaceshipFragment)
	fs.Compile()

	return fs
}
