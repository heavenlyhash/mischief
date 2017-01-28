http://adriansampson.net/blog/opengl.html

- dot products, FOVs, other practicalities explained
  - http://blog.wolfire.com/2009/07/linear-algebra-for-game-developers-part-2/

A nice series of write-ups with specific focus on voxels in practice:
  https://sites.google.com/site/letsmakeavoxelengine/

Another golang voxel project:
  - https://github.com/boombuler/voxel/tree/master/rendering
  - has mesh-generation (and voxel culling) in meshing.go
    - end result is a '[]VertexF' (it's {Color,Norm,Pos}) which pretty much ready to flip into a VBO
  - has frustrum culling
  - has interleaved arrays (see meshbuffer.go)
    - interesting use of `gl.InterleavedArrays(gl.C4F_N3F_V3F [...])`.  Seems to elide need for other VAO?

Another golang voxel project:
  - https://github.com/allanks/Voxel-Engine/
  - pretty large amount of player-move-in-gravity code in src/Player/Player.go
  - straightforward texture image repacker in src/TexturePacker/TexturePacker.go
  - otherwise... pretty messy; loops and networks in the middle of the terrain code, oh my
  - use of bitmap fonts, plus the convert-ttf-to-bmp part, can be found in src/glText45
  - parser for "obj" format, which appears to be pretty standard: src/ObjectLoader/ObjectLoader.go
    - yields {vertices, normals, uv/texture}
  - A more complex approach than most! src/Graphics/Game/OpenGL45/OpenGL45.go

Not very interesting projects:
  - https://github.com/samnm/goblocks/ -- not far along; only every uses one program.

A wise person's attempt at GL helper wrappers:
  - https://godoc.org/github.com/go-gl-legacy/glh
  - https://github.com/go-gl-legacy/glh/blob/master/meshattr.go#L26
  - i'm unsure why this is considered deprecated; it looks none too crazy.
