package game

import (
	"fmt"

	"github.com/morcmarc/gosteroids/game/objects"
	"github.com/morcmarc/gosteroids/game/shaders"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
)

var (
	vArray  gl.VertexArray
	vBuffer gl.Buffer
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

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.MULTISAMPLE)
	gl.Disable(gl.LIGHTING)
	gl.Enable(gl.COLOR_MATERIAL)

	fmt.Println(":: OpenGL Context initialized.")

	vArray = gl.GenVertexArray()
	vArray.Bind()

	vBuffer = gl.GenBuffer()
	vBuffer.Bind(gl.ARRAY_BUFFER)

	program := bindShaders()
	fmt.Println(":: Compiled shaders.")

	positionAttrib := program.GetAttribLocation("position")
	positionAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)
	positionAttrib.EnableArray()

	mainLoop(window, ctrlChnl)

	defer positionAttrib.DisableArray()
	defer window.Destroy()
	defer glfw.Terminate()
	defer program.Delete()
	defer vBuffer.Delete()
}

func mainLoop(window *glfw.Window, ctrlChnl chan uint8) {
	fmt.Println(":: Starting loop.")

	sship := objects.NewSpaceship()
	bg := objects.NewBackground()

	for !window.ShouldClose() {
		drawScene(bg, sship)

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
	shaders := shaders.CompileAll()

	for _, shader := range shaders {
		program.AttachShader(shader)
		defer shader.Delete()
	}

	program.BindFragDataLocation(0, "outColor")
	program.Link()
	program.Use()

	return program
}

func drawScene(bg *objects.Background, sship *objects.Spaceship) {
	gl.ClearColor(0.1, 0.1, 0.2, 0.5)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	bg.Draw()
	sship.Draw()
}

func errorCallback(err glfw.ErrorCode, desc string) {
	if glerr := gl.GetError(); glerr != gl.NO_ERROR {
		string, _ := glu.ErrorString(glerr)
		panic(string)
	}
}
