package graphics

import (
	"github.com/go-gl/gl"
)

func NewProgram(vs gl.Shader, fs gl.Shader) gl.Program {
	program := gl.CreateProgram()
	program.AttachShader(vs)
	program.AttachShader(fs)
	program.Link()

	return program
}
