package thunk

/*
	Bunches up a sequence of drawing operations.

	This is a fluent interface meant to make it easy to compose chunks of
	rendering info on one thread, then toss it to the actual renderer, who
	commits it as fast as possible onto the GPU.
	Some things are still presumed to be batched (e.g. camera, perspective,
	whatever your program wants -- the "uniforms" in GL land); vector
	arrays and buffers are expected to come in whole streams per set
	of uniform state.
*/

// configure the program
//   - setting uniforms.
//        **this is a program-specific thing**, with no general values
//   - pipe arrays and buffers, and hit draw

// you have to have a type matching your program and all its uniforms

type Thunk interface {
	Start() // you should: call `useProgram`, commit your uniforms.
	// this is weird... the model coord matrix uniform is frequently changing, but the rest (camera, etc) is not
	// so clearly changing uniforms isn't that expensive either
	// however, i defintely don't want ever object in the game to care abou the cemera so
	// maybe this should return

	// REGARDLESS, this is more to do with chunks and shouldn't be in the `glo` pkg

	Itr() <-chan DrawItem
}

type DrawItem struct {
	// i think this can be just vertex stuff with no magic?
	// gl 'attribute's are per vertex, that's essentially what this is
	// welp i guess those are still an interface thing too
	// there may be a whole bunch of vertexes per attribute batch one hopes
	//	/Start() // you should: commit your attributes
	// TODO getter for vertexes?  those buffers might need to be named&alloc'd from the gl thread...
}

// i want
//   - a way to make sure i specified all program inputs, structurally
//     - this is an inherently weird desire, because structures spec some of these, but not most: camera?  no.
//   - a way for major objects (or terrain chunks) in the game to define their own rendering
//     - these can internally try to sequence operations for min texture swapping etc if want (chunks do culling, etc)
//     - do these have more than one set of vbo's inside?  YES.  chunks don't; fancy machines and pipes probably do.
//   - a way to cache vao/vbo's non-redundantly?
//     - i'm not sure there's a general-purpose way to help with this
//   - the renderer needs a general way to tell if something's within a million miles or not
//     - it can then give that thing a strong hint it's not going to be needed soon
//     - this reminds me of LOD levels
