package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"exultant.us/mischief/cmd/mischief/controls"
	"exultant.us/mischief/lib/maath"
	"exultant.us/mischief/mirage"
	"exultant.us/mischief/render/prag"
	"exultant.us/mischief/render/texture"
)

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	runtime.LockOSThread()

	viewport := maath.Vec2i{800, 600}

	// SO MUCH SETUP.  Start with setting up window and GL pre-init params.
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DepthBits, 32)
	glfw.WindowHint(glfw.StencilBits, 0)
	window, err := glfw.CreateWindow(viewport.X(), viewport.Y(), "Mischief", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1) // vsync
	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}
	// Configure yet more global GL settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.01, 0.06, 0.03, 1.0)

	// Load textures.
	// Best to do this up front so we don't have random stalls later as things
	//  discover they need textures during their first render.  (This has GL
	//  calls, but happily program-agnostic.)
	tCache := texture.NewCache()
	tCache.Load("placeholder", "assets/texture/placeholder.png")
	checkGLError()

	// Make a program.
	prog := mirage.NewBasicProgram()
	checkGLError()
	// Tell it to preload all its stuff.
	prog.Preload()
	checkGLError()

	// Make camera.
	cam := &controls.Camera{}
	cam.Drifter.InitDefaults(mgl32.Vec3{2, 2, -3})
	// Grab cursor.  Route to camera.
	window.SetCursorPosCallback(
		func(window *glfw.Window, xpos, ypos float64) {
			// Also forces resetting the cursor to the center of the screen,
			//  which is the only reason this is here instead of hidden in the package with the camera as yet.
			center := viewport.Vec2().Mul(.5)
			window.SetCursorPos(float64(center[0]), float64(center[1]))
			cam.Rotate(mgl32.Vec2{
				float32(ypos) - center[1],
				float32(xpos) - center[0],
			})
		},
	)

	// Make a thing!
	obj := &mirage.Cube{}

	// Polllllll
	for !window.ShouldClose() {
		checkGLError()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Move camera
		controls.UpdateCameraFromKeyboard(cam, window)
		cam.Tick()

		// Tell the program it's in charge.
		// Update the major uniforms (e.g. camera).
		prog.Arrange()
		projectionMtrx := mgl32.Perspective(mgl32.DegToRad(75.0), float32(viewport.X())/float32(viewport.Y()), 0.1, 50.0)
		prog.(prag.ProgramWithProjection).SetProjection(projectionMtrx)
		checkGLError()
		prog.(prag.ProgramWithCamera).SetLook(cam.GetLookMatrix())
		checkGLError()

		// Render
		// (model coords still hardcode, fixme shortly)
		prog.SetModel(mgl32.Ident4())
		obj.Render(tCache)
		checkGLError()

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func checkGLError() {
	err := gl.GetError()
	if err != 0 {
		panic(fmt.Sprintf("gl error: %d", err))
	}
}
