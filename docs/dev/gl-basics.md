what is opengl, really
======================

Windows, Programs, Shaders (for vertex, fragment, and geometry), oh my!

Also Framebuffers.  Framebuffers are for rendering to a texture, e.g. for making a PIP or such.
(See http://www.opengl-tutorial.org/intermediate-tutorials/tutorial-14-render-to-texture/ for more about that.)

Programs are the level where most of the call interface happens.
A render call requires:
- Set the program
- Set the uniforms
- Set the attributes
- Set the VAO
- Set the VBO
- Call draw.



what are GL versions
--------------------

GL versioning is mildly terrifying.

We're chosing 4.3.  Reasons:

- 4.3 introduced debugging calls which are significantly useful.
- it looks like it should still be on the safe side for compat after 2012 certainly (and anecdotally, I'd say farther).

Other scare factors for GL version availability:
  - Mac OS X Mavericks is limited to OpenGL 4.1
    - supplanted late 2014; released late 2013; still supported in 2016
	  - no idea what market share is but this is after the beginning of their free upgrades policy so likely low.



the egl Program API
-------------------

So where do we put all these things together?

A couple groupings:

- **Programs** -- the validity of uniform and attribute calls is defined by what program is active, so it's a necessary scoper.
- **Mangles** -- one batch of attributes, vao/vbo's, and draw calls.  These tend to have types (e.g. a "cube" or a "chunk" -- all their instances have the same attributes, just different volumes of data).  Mangles can theoretically be shared across programs, but the attributes must match up, so in practice they tend not to.
- **Framebuffers** -- when you start invoking a draw, you must have one bound.

Efficiency concerns:

- If you can use a program consistently without switching, it saves a bunch of uniform setups and such.
- If you can update a subset of the uniforms for each new Mangle and draw call batch, you save a few calls (e.g. in practice usually only the model coord transform changes).
- If you can call a batch of same-type Mangles (instead of varying types interleaved), you can save a bunch of attribute bind calls.
- If you can memoize your VBOs... well, obviously, do so for any nontrival ones.
- Texture binding is somewhere intermixed with the draw functions in a Mangle.  It's problematic to try for much efficiency here,
  since we're using Mangles as a high-level grouping of objs that should or should not be rendered.  An intermediate layer
  which buffered up tons of {tex,vao,vbo,drawfn} tuples would be able to sort these for efficiency, but adds a lot of complexity.
  Using `BindTexture(gl.TEXTURE_2D_ARRAY,...)` may buy a lot of mileage without much complexity: with this, you can design
  such that there's simply need for more than one texture per Mangle, thus making mere Mangle-type-batching DTRT for textures as well.

So in terms of what happens within which scope limits:

- Defining a Program/Shader source is a (practically speaking) compile-time thing.
- Defining a Mangle's attributes and draw calls is a (practically speaking) compile-time thing.
- Computing a Mangle's innards (e.g. memoizing the VBO contents) is scope-free.
- Render calls should happen in the following nested scope:
  - Window and Init
    - Bind Framebuffer
	  - Bind Program
	    - Call `program.Draw(mangle)` -- which binds attribs and vbo and invokes the draw thunks

### things we're choosing to ignore

You can render to several framebuffers at the same time.
It's pretty hard for me to imagine a good use for this unless it's for scaling an image or something.
It also requires adding more outputs to your fragment shader, so a very broad reach is required to do this.
(It seems to maybe come up in shadowmaps?  But at present I don't deeply understand this.)


```
program := CompileProgram(Shader{"src", ShaderT_Vertex})
program.Bind("camera", mat4{})
program.Bind("perspective", mat4{})
```

