/*
	Like a prog, but more pragmatic.

	In seriousness though.
	Wrapper around GL "programs" so you spend less time futzing around with the
	program handle and more time just giving it a dang byte range.
*/
package prag

import (
	"github.com/go-gl/mathgl/mgl32"
)

/*
{{{ general render setup phase
- set program
- bind vao
- enable attributes
- set attribute offsets
- set common uniforms
  - camera, perspective, these seem to be thing-independent
    - mayb wise to have our structure give these well-known interfaces if it wants to
}}}
{{{ per-thunk self-directed render phase
- !!! set some uniforms
  - some question here if we want to treat model coords separately?
- !!! possiblely some texture shuffling
  - though afiak we might want to use serious atlases to nearly never do this so who knows what that api looks like
- !!! do as much byte shoveling as possible
  - so we want to tie a struct back to the "set attribute offsets" part
    - this is Not Possible to do in a generic type safe interface in go
	  - unless done very carefully with the sort-style inversion of interface
  - Things should cache their vbo automatically.
}}}
*/

type Program interface {
	SetModel(mgl32.Mat4) // we consider a "where" matrix too common to be optional
	Preload()            // called once at creation.  set up your textures, etc.  (grabbag, seemed needed)
	Arrange()            // called pre-draw group.  bind vao, set all attribs up, etc.  (effectively, magic stuff and offsets that your RenderableThing's data layout must agree with go in here.)

	// Try to specify any additional interesting uniforms via interfaces.
	//  See e.g. `ProgramWithCamera` for an example.
	//  Or, put hide them inside `Preload()` or `Arrange()`.
	//  Use a non-standard method that requires an explicit cast to your program type as a last resort.textex
}

type ProgramWithCamera interface {
	SetLook(mgl32.Mat4)
}

type ProgramWithProjection interface {
	SetProjection(mgl32.Mat4)
}

/*
	The thing itself should be something we can toss as a `gl.Ptr` (generally a
	slice, where each instances of the type describes all the attributes per vertex).
	The type's fields inherently need to line up with the program description; that
	part's your problem.

	FIXME : you probably need a RenderPlanner step before that, so it can decide
	whether or not to bother producing the whole list of vertex attribs!!
	I guess that means the interface will have to return an interface{}, sigh ok.
*/
type RenderableThing interface {
	RendersWith() Program
	NeedsRerender() bool  // if true, blow any cached vbo and do everything fresh
	Position() mgl32.Mat4 // where to start the render at (coords & rot), may change freely even if vbo doesn't need rerendering
}
