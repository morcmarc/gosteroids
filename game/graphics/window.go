package graphics

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	. "github.com/morcmarc/gosteroids/game/shared"
)

var (
	controlChanel chan uint8
	ErrGLFWFailed = errors.New("Failed to init glfw")
)

func Init(width, height int, title string, ctrlChnl chan uint8) {
	controlChanel = ctrlChnl

	window, err := initGL(width, height, title)
	if err != nil {
		log.Printf("InitGL: %v", err)
		return
	}
	defer glfw.Terminate()

	scene := NewScene()

	for !window.ShouldClose() {
		// Reset
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.LoadIdentity()

		scene.Draw()

		// Render
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

// initGL initializes GLFW and OpenGL.
func initGL(width, height int, title string) (*glfw.Window, error) {
	ok := glfw.Init()
	if !ok {
		return nil, ErrGLFWFailed
	}
	glfw.SwapInterval(1)

	// glfw.WindowHint(glfw.ContextVersionMajor, 3)
	// glfw.WindowHint(glfw.ContextVersionMinor, 3)
	// glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	// glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, err
	}

	window.SetKeyCallback(onKey)
	window.MakeContextCurrent()

	gl.Init()
	if err = glh.CheckGLError(); err != nil {
		panic(err)
	}

	renderer := gl.GetString(gl.RENDERER)
	version := gl.GetString(gl.VERSION)
	fmt.Printf("Renderer: %s\n", renderer)
	fmt.Printf("OpenGL version supported: %s\n", version)

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.TEXTURE_2D)
	gl.Disable(gl.LIGHTING)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	return window, nil
}

// onKey handles key events.
func onKey(window *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, _ glfw.ModifierKey) {
	if key == glfw.KeyEscape {
		window.SetShouldClose(true)
	}

	if key == glfw.KeyUp {
		controlChanel <- Throttle
	}

	if key == glfw.KeyDown {
		controlChanel <- Break
	}

	if key == glfw.KeyLeft {
		controlChanel <- Left
	}

	if key == glfw.KeyRight {
		controlChanel <- Right
	}
}
