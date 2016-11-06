package eglProgram

/*
	Shader is a ref to a segment of code compiled onto the GPU.
	String them up into `Program`s, then tell the GPU to use that program
	when rendering.

	We don't export this because there's just not a whole ton of call for
	reusing shaders between different programs, so for simplicity
	our API just combines the compile steps of both.

	REVIEW: worth considering exporting a separate compile step just for
	getting that hair more stack info in the case of compile errors.
*/
type shader uint32
