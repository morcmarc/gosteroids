package game

import (
	"github.com/morcmarc/gosteroids/game/shaders"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
)

func InitWindow(width, height int, title string, ctrlChnl chan uint8) {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Could not init glfw")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	vao := gl.GenVertexArray()
	vao.Bind()

	vbo := gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)

	verticies := []float32{0, 1, 0, -1, -1, 0, 1, -1, 0}

	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, verticies, gl.STATIC_DRAW)

	program := bindShaders()

	positionAttrib := program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()

	mainLoop(window, ctrlChnl)

	defer positionAttrib.DisableArray()
	defer window.Destroy()
	defer glfw.Terminate()
	defer program.Delete()
}

func mainLoop(window *glfw.Window, ctrlChnl chan uint8) {
	for !window.ShouldClose() {
		drawScene()

		window.SwapBuffers()
		glfw.PollEvents()

		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}

		if window.GetKey(glfw.KeyUp) == glfw.Press {
			ctrlChnl <- THROTTLE
		}

		if window.GetKey(glfw.KeyDown) == glfw.Press {
			ctrlChnl <- BREAK
		}

		if window.GetKey(glfw.KeyLeft) == glfw.Press {
			ctrlChnl <- LEFT
		}

		if window.GetKey(glfw.KeyRight) == glfw.Press {
			ctrlChnl <- RIGHT
		}
	}
}

func bindShaders() gl.Program {
	program := gl.CreateProgram()
	program.AttachShader(shaders.GetSpaceshipVertex())
	program.AttachShader(shaders.GetSpaceshipFragment())

	program.BindFragDataLocation(0, "outColor")
	program.Link()
	program.Use()

	return program
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func errorCallback(err glfw.ErrorCode, desc string) {
	if glerr := gl.GetError(); glerr != gl.NO_ERROR {
		string, _ := glu.ErrorString(glerr)
		panic(string)
	}
}
