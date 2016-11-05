/*
	The window package wraps up all the calls to `glfw` necessary to create a root
	window for us to draw in.

	It makes many assumptions to simplify this:

	  - the GL versions are chosen for you -- 3.2.
	  - you're getting vsync.
	  - the initial size is currently hardcoded
	    (we just make sure resize works reliably).

	We haven't made any attempt to DTRT with more than one window.
*/
package eglWindow

import "github.com/go-gl/glfw/v3.2/glfw"

type Window struct {
	window *glfw.Window
}

func (w *Window) Start() {
	staticInitializers()
	var err error
	w.window, err = glfw.CreateWindow(800, 600, "Mischief", nil, nil)
	if err != nil {
		panic(err)
	}
	w.window.MakeContextCurrent()
}

func staticInitializers() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	// Set GL pre-init params.  Globals(?).
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DepthBits, 32)
	glfw.WindowHint(glfw.StencilBits, 0)
	glfw.SwapInterval(1) // vsync
}

// Flip buffers and handle any buffered events.
// Call at the end of every draw cycle.
func (w *Window) StepMaintenance() {
	w.window.SwapBuffers()
	glfw.PollEvents()
}

func (w *Window) Terminate() {
	glfw.Terminate()
}
