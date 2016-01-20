package mirage

import (
	"github.com/go-gl/gl/v3.2-core/gl"

	"exultant.us/mischief/render/texture"
	"exultant.us/mischief/render/vertex"
)

/*
	Ze Cube Du Debug
*/
type Cube struct {
}

func (obj *Cube) Render(tCache *texture.Cache) {
	// Push cube vertex data into the buffer.
	// Meanings and layouts were already configured for us in the program's `Arrange()` cycle.
	gl.BufferData(gl.ARRAY_BUFFER, len(vertex.Cube)*4, gl.Ptr(vertex.Cube), gl.STATIC_DRAW)

	// Set our texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tCache.Get("placeholder"))

	// Do draw!
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
	gl.GetError()
}
