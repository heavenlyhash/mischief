package main

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DepthBits, 32)
	glfw.WindowHint(glfw.StencilBits, 0)

	window, err := glfw.CreateWindow(854, 480, "Mischief", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1) // vsync

	gl.Init()
	// TODO you can load textures after this

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
