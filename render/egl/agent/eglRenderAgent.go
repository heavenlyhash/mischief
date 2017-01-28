package eglAgent

import (
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"

	"exultant.us/mischief/render/egl/window"
)

func New(wiring Wiring) *Renderer {
	return &Renderer{
		wiring: wiring,
	}
}

type Renderer struct {
	state  state
	wiring Wiring
}

type Wiring struct {
	callNextFrame <-chan func()
}

type state struct {
	window *eglWindow.Window
}

/*
	Runs, pinning its goroutine to an OS thread, and performing *all* the
	operations for GL systems (as essentially nothing about GL is multi-thread safe).

	This means there's a *lot* of setup for this actor (since it does *everything*,
	including setting up the very first window), and most of the behaviors
	thereafter come from communication via channels (and sometimes callbacks).
*/
func (a *Renderer) Run() {
	a.state.window = &eglWindow.Window{}
	a.state.window.Start()
	for {
		a.step()
	}
}

func (a *Renderer) step() {
	// Raise any errors.
	mustCheckGLError()

	// Clear screen (appropriate to just redo in a FPS-style always-changing view).
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Call arbitrary funcs that were pushed into our todo list.
	// (Texture loading currently flies through here, for example.)
	for {
		callme, ok := <-a.wiring.callNextFrame
		if !ok {
			break
		}
		callme()
	}

	// Render
	// TODO thunks go here

	// Maintenance
	a.state.window.StepMaintenance()
}

func mustCheckGLError() {
	err := gl.GetError()
	if err != 0 {
		panic(fmt.Errorf("gl error: %d", err))
	}
}
