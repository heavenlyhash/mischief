package mirage

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"exultant.us/mischief/render/vertex"
)

/*
	Ze Cube Du Debug
*/
type Cube struct {
	Coord mgl32.Vec3
}

func (obj *Cube) Render(programID uint32) (vao, vbo uint32) {
	// Ask GL for handles to arrays.
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	// Turn around and tell GL that we're going to use that array.
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Push cube vertex data into the buffer.
	gl.BufferData(gl.ARRAY_BUFFER, len(vertex.Cube)*4, gl.Ptr(vertex.Cube), gl.STATIC_DRAW)

	// Enable the 'attributes' to the GL program.
	// They're not enabled by default because GL loves to be redundant.
	// Also we need to be able to refer to these 'attributes' so that
	//  we can turn around and tell GL where to get their values from.
	vertAttrib := uint32(gl.GetAttribLocation(programID, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	texCoordAttrib := uint32(gl.GetAttribLocation(programID, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)

	// Tell GL how to interpret locations in our arrays onto meanings in the program.
	// - name the attribute to assign from slot
	// - specify how many there are (e.g. 3 for a vec3)
	// - specify the type
	// - ??? backcompat flag that's always false afaict
	// - "stride" length
	// - offset per stride
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	// Do draw!
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

	return
}
