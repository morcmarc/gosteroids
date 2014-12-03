package graphics

import (
	"io/ioutil"

	"github.com/go-gl/gl"
)

type ShaderType uint8

const (
	VertexShader   ShaderType = 0
	FragmentShader ShaderType = 1
)

func LoadShader(path string, shaderType ShaderType) (gl.Shader, error) {
	var shader gl.Shader

	sourceByte, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}
	source := string(sourceByte)

	switch shaderType {
	case VertexShader:
		shader = gl.CreateShader(gl.VERTEX_SHADER)
		break
	case FragmentShader:
		shader = gl.CreateShader(gl.FRAGMENT_SHADER)
		break
	}

	shader.Source(source)
	shader.Compile()

	return shader, nil
}
