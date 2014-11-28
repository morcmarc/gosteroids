package shaders

import (
	"github.com/go-gl/gl"
)

func CompileAll() []gl.Shader {
	ssv := CompileVertexShader(SpaceshipVertex)
	ssf := CompileFragmentShader(SpaceshipFragment)

	sl := []gl.Shader{
		ssv,
		ssf,
	}

	return sl
}

func CompileVertexShader(s string) gl.Shader {
	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vs.Source(s)
	vs.Compile()

	return vs
}

func CompileFragmentShader(s string) gl.Shader {
	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fs.Source(s)
	fs.Compile()

	return fs
}
