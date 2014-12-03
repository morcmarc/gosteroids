package graphics

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	o "github.com/morcmarc/gosteroids/game/objects"
	. "github.com/morcmarc/gosteroids/game/shared"
)

var (
	controlChanel chan uint8
	ticker        *time.Ticker
	hasTicked     bool
	currentTime   float32 = 0.0
	objectManager *o.ObjectManager
)

func Init(width, height int, title string, cc chan uint8, om *o.ObjectManager) {
	controlChanel = cc
	objectManager = om

	window, err := initGL(width, height, title)
	if err != nil {
		log.Printf("InitGL: %v", err)
		return
	}
	defer glfw.Terminate()

	// Set up ticker to 60 FPS
	ticker = time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	hasTicked = false

	go func() {
		for _ = range ticker.C {
			currentTime += 0.008
			hasTicked = true
		}
	}()

	scene := NewScene(objectManager)

	for !window.ShouldClose() {
		// Reset
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)

		// Update state?
		if hasTicked {
			glfw.PollEvents()
			checkKeys(window, controlChanel)
			scene.Update(currentTime)
			hasTicked = false
		}

		// Render
		scene.Draw(currentTime)
		window.SwapBuffers()
	}
}

// initGL initializes GLFW and OpenGL.
func initGL(width, height int, title string) (*glfw.Window, error) {
	ok := glfw.Init()
	if !ok {
		return nil, errors.New("Failed to init glfw")
	}
	glfw.SwapInterval(1)

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, err
	}

	// window.SetInputMode(glfw.StickyKeys, 1)
	window.MakeContextCurrent()

	gl.Init()
	if err = glh.CheckGLError(); err != nil {
		// panic(err)
	}

	renderer := gl.GetString(gl.RENDERER)
	version := gl.GetString(gl.VERSION)
	fmt.Printf("Renderer: %s\n", renderer)
	fmt.Printf("OpenGL version supported: %s\n", version)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	return window, nil
}

func checkKeys(window *glfw.Window, cc chan uint8) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}

	u := window.GetKey(glfw.KeyUp)
	d := window.GetKey(glfw.KeyDown)
	l := window.GetKey(glfw.KeyLeft)
	r := window.GetKey(glfw.KeyRight)

	if l == glfw.Press {
		cc <- Left
	}

	if r == glfw.Press {
		cc <- Right
	}

	if u == glfw.Press {
		cc <- Throttle
	}

	if d == glfw.Press {
		cc <- Break
	}
}
