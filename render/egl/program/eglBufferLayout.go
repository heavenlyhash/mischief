package eglProgram

type BufferLayout []BufferAttribute

type BufferAttribute struct {
	Name       string      // Used when mapping into the `Program`.  Irrelevant during buffer filling.
	Size       int         // The number of components in the attribute (e.g. 3 for a vec3).
	Type       interface{} // Currently ignored and always "float".
	Normalized bool        // Almost always false.

	// Most of these params are as per here in the GL spec:
	//   https://www.opengl.org/sdk/docs/man/html/glVertexAttribPointer.xhtml

	// We compute the following attributes for you instead,
	//  which is easy to do within our api since we've declared that your
	//  attributes are always fully packed with no other intervening data,
	//  thus disallowing any of the complexity which the customizable stride
	//  and jump parameters were designed to permit.

	stride int // How wide each full group of attributes is, in bytes.
	jump   int // How far into the stride this attribute starts, in bytes.
}

// Ah.  You *do* have to have a *particular* buffer bound before calling the attribthingy func.
// You *can* have separate buffers.
// Kinda sounds like "which buffer" is also stored in the VAO.
