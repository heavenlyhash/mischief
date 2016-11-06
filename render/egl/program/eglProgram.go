package eglProgram

/*

This package exposes the basic functions to compile and use GL programs.
You're going to want to create your own types to represent your *particular*
GL program and its parameters with real types -- just use this interface
to help sanity check things at load time and help you with type conversion.

## How to define Programs:

You *don't* need whole different programs just to draw different kinds of geometry:
that's all in your vertex buffers.
You *do* need different programs when you want different kinds of drawing logic:
any time you have different texture mapping choices, color ranging tricks, etc,
then you're talking about things that require shader logic.

In a first-person perspective, almost every program is going to have the same:

  - cameraMatrix
  - positionMatrix
  - perspectiveMatrix

A program that's drawing part of the world usually as an offset matrix for the
thing you're drawing.

## What Program usage looks like:

GL is all about getting things lined up in batches, so the GPU can parallelize them.

So, you tend to want/need to call things in a sequence like this:

  - Tell the GPU what program to use to understand the next few calls.
  - Set some "uniforms".  These are global variables for the next few calls.
  - Set "attributes" configuration.  This will tell the GPU how to view the next value.
  - Foist in a "buffer".  The buffer will be divide-n-conquer'd in accordance with the attributes cfg.
  - Call draw.  (The GPU will act on each row of the attributes in parallel.)
  - GOTO foist-a-buffer and repeat!
  - GOTO set-some-uniforms and repeat!

Anything that you can skip pushing again between draw calls because it hasn't changed will
improve your overall performance by avoiding the time taken to round-trip communicating with the GPU.

Picking what to skip is complicated: the correct answer again depends on your application.

The buffer generally changes on every draw call.
(It may be worth caching some buffers in your application logic (for example,
the vertexes describing your terrain probably don't change that often).)

The attributes configuration changes nearly never.
(The wild amount of configurability possible for "strides" etc is mostly
from the legacy of C code, where you might be storing a bunch of vertex info
mixed into an array of other logical objects.  We tend to simply not do that.)

Uniforms generally change much less than once per draw call.
For example, most FPS-style games will have those camera and perspective uniforms
be essentially constant for a huge number of rounds of varying attribute vectors.
But not all uniforms are the same: the model position offset uniform changes,
often along with each vertex buffer, in the example of terrain chunk draws.

So, we offer full flexibility.

You define all the things needed for a full render call when you construct your Program.

Thereafter, you bind named values, which returns a function that's essentially
a partial evaluation.
Binding to a name twice panics, because it's a programmer error;
attempting to call the final draw without having bound each value also panics.
Thus, you can hold onto the partial evaluations as long as their bindings are
valid, and the API helps you make sure you called everything, no more and no less.

*/

func CompileProgram(shaders ...struct {
	shaderMode int
	shaderSrc  string
}) *Program {
	p := &Program{}
	p.link()
	p.initVAO()
	return p
}

type Program struct {
	glHandle     uint32
	vao_glHandle uint32
	uniforms     []struct {
		Name     string
		Setter   func(interface{})
		glHandle uint32
	}
}

func (p *Program) link() {
}
func (p *Program) initVAO() {
	// REVIEW: we may want to make these part of the same API as buffer fabbing,
	//  because something needs to make filling those arrays make any goddamn sense too.
	gl.GenVertexArrays(1, &p.vao_glHandle) // Ask for VAO allocation.
	gl.BindVertexArray(p.vao)              // Turn around and bind it.
}

func (p *Program) Bind(fieldName string, yourState interface{}, nextFn func(*BoundProgram)) {

}

type BoundProgram struct {
	p       *Program
	isBound []bool
}

func (bp *BoundProgram) Ready() bool {
	return true
}

/*
	The terminal thunk.
*/
func Draw(bp *BoundProgram) {
	// TODO might still need the buffer contents to be pushed:
	/*
		gl.BufferData(gl.ARRAY_BUFFER, len(cube_vertices)*4, gl.Ptr(cube_vertices), gl.STATIC_DRAW)
	*/
	// TODO might still need textures to be bound for the call:
	/*
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tCache.Get("placeholder"))
	*/
	// TODO draw calls are themselves still specialized and param'd:
	/*
		gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	*/
}
