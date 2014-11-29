package game

import (
	"fmt"

	"github.com/morcmarc/gosteroids/game/objects"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glu"
)

func InitWindow(width, height int, title string, ctrlChnl chan uint8) {
	glfw.SetErrorCallback(errorCallback)

	if !glfw.Init() {
		panic("Could not init glfw")
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	gl.Enable(gl.TEXTURE_2D)

	fmt.Println(":: OpenGL Context initialized.")

	mainLoop(window, ctrlChnl)

	defer window.Destroy()
	defer glfw.Terminate()
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

func drawScene(bg *objects.Background, sship *objects.Spaceship) {
	gl.ClearColor(0.1, 0.1, 0.2, 0.5)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	bg.Draw()
}

func errorCallback(err glfw.ErrorCode, desc string) {
	if glerr := gl.GetError(); glerr != gl.NO_ERROR {
		string, _ := glu.ErrorString(glerr)
		panic(string)
	}
}
