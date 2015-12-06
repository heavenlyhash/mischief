package main

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"exultant.us/mischief/render"
	"exultant.us/mischief/render/shader"
	"exultant.us/mischief/render/texture"
	"exultant.us/mischief/render/vertex"
)

// fuckkit
const windowWidth = 800
const windowHeight = 600

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

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// ??? WIZARDRY
	programID := render.NewProgram(
		shader.PlaceholderVertexShader,
		shader.PlaceholderFragmentShader,
	)
	gl.UseProgram(programID)

	// Load a texture.
	txture := texture.FromFile("assets/texture/placeholder.png")

	// Okay, start hucking things on the screen

	projectionMat := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 10.0)
	projectionID := gl.GetUniformLocation(programID, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionID, 1, false, &projectionMat[0])

	cameraMat := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cameraID := gl.GetUniformLocation(programID, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraID, 1, false, &cameraMat[0])

	modelMat := mgl32.Ident4()
	modelID := gl.GetUniformLocation(programID, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelID, 1, false, &modelMat[0])

	textureID := gl.GetUniformLocation(programID, gl.Str("tex\x00"))
	gl.Uniform1i(textureID, 0) // ?

	gl.BindFragDataLocation(programID, 0, gl.Str("outputColor\x00"))

	// Configure the vertex data
	// (Erics are somewhat confused by this.  It seems to be poking
	//  horifficially global variables in video memory by string name...?)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertex.Cube)*4, gl.Ptr(vertex.Cube), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(programID, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(programID, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.01, 0.06, 0.03, 1.0)

	// Polllllll
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Render
		gl.UseProgram(programID)
		gl.UniformMatrix4fv(modelID, 1, false, &modelMat[0])

		gl.BindVertexArray(vao)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, txture)

		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
