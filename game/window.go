package game

import (
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
)

func InitWindow(w, h int, t string, ctrlChnl chan uint8) {
	runtime.LockOSThread()

	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Could not init glfw")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(w, h, t, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()
	initScene()

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

	defer window.Destroy()
	defer glfw.Terminate()
}

func initScene() {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
}

func drawScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func errorCallback(err glfw.ErrorCode, desc string) {
	if glerr := gl.GetError(); glerr != gl.NO_ERROR {
		string, _ := glu.ErrorString(glerr)
		panic(string)
	}
}
