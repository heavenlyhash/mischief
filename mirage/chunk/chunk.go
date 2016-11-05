package chunk

import (
	"encoding/binary"
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"

	"exultant.us/mischief/matter/direction"
	"exultant.us/mischief/render/texture"
)

type BlockType string

const blocksPerChunk = 16 * 16 * 16

// 16x16x16 hunks.  because offsets for those are pleasant to calculate.
type Model struct { // Move to game logic latur
	blocks [blocksPerChunk]BlockType
}

func (m *Model) blockIndex(x, y, z int) int {
	return (y << 8) | (z << 4) | x
}

func (m *Model) GetBlockAt(x, y, z int) BlockType {
	return m.blocks[m.blockIndex(x, y, z)]
}

func (m *Model) SetBlockAt(x, y, z int, bt BlockType) {
	m.blocks[m.blockIndex(x, y, z)] = bt
}

type Mirage struct {
}

func (obj *Model) Render(tCache *texture.Cache) {
	nVerts := 0
	vertbuf := make([]chunkVertex, 0)
	for x := 0; x < 1; x++ {
		for z := 0; z < 1; z++ {
			for y := 0; y < 1; y++ {
				//ib := obj.blockIndex(x, y, z)
				//fmt.Printf("*> %d %d %d\n", x, y, z)
				for _, face := range direction.Faces {
					switch face {
					case direction.North:
					case direction.East:
					default:
						continue
					}
					for _, vert := range faceVertices[face].verts {
						nVerts++
						v := vert
						v.X += float32(x)
						v.Y += float32(y)
						v.Z += float32(z)
						//fmt.Printf("v>>   pts:%2.0f,%2.0f,%2.0f\n", v.X, v.Y, v.Z)
						vertbuf = append(vertbuf, v)
					}
				}
			}
		}
	}

	// Push cube vertex data into the buffer.
	// Meanings and layouts were already configured for us in the program's `Arrange()` cycle.
	gl.BufferData(gl.ARRAY_BUFFER, nVerts*vertStride, gl.Ptr(vertbuf), gl.STATIC_DRAW)

	// Set our texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tCache.Get("placeholder"))

	// Tell GPU about the vertex repeat sauce
	//	var veb uint32
	//	gl.GenBuffers(1, &veb)
	//	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, veb)
	//	ebuff := genElementBuffer(nVerts)
	//	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(ebuff), gl.Ptr(ebuff), gl.DYNAMIC_DRAW)

	// Do draw!
	fmt.Printf("d>>   verts:%4d\n", nVerts)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(nVerts))

	//	gl.DeleteBuffers(1, &veb)
}

func genElementBuffer(size int) []byte {
	data := make([]byte, size*4)
	offset := 0
	//gl.UnsignedInt
	batches := size / 6
	for i := 0; i < batches; i++ {
		for _, val := range []uint32{0, 1, 2, 2, 1, 3} { // why the fuck is this different than the facedetails one
			binary.LittleEndian.PutUint32(data[offset:], uint32(i)*4+val)
			offset += 4
		}
	}
	return data
}

const vertStride = 12 + 8 + 6 + 2 + 3 + 1 + 4 + 4

type chunkVertex struct {
	X, Y, Z                    float32
	TX, TY, TW, TH             uint16 // ?
	TOffsetX, TOffsetY, TAtlas int16
	Pad0                       int16
	R, G, B                    byte
	Pad1                       byte   // x
	BlockLight, SkyLight       uint16 // x
	Pad2, Pad3                 uint16 // x
}

type faceDetails struct {
	indices [6]int32
	verts   [6]chunkVertex
}

// Precomputed face vertices
var faceVertices = [6]faceDetails{
	{ // North
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 0, Y: 0, Z: 0, TOffsetX: 1, TOffsetY: 1},
			{X: 1, Y: 0, Z: 0, TOffsetX: 0, TOffsetY: 1},
			{X: 0, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 0, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 0, Y: 0, Z: 0, TOffsetX: 1, TOffsetY: 1},
			{X: 1, Y: 1, Z: 0, TOffsetX: 0, TOffsetY: 0},
		},
	},
	{ // East
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 1, Y: 0, Z: 0, TOffsetX: 1, TOffsetY: 1},
			{X: 1, Y: 0, Z: 1, TOffsetX: 0, TOffsetY: 1},
			{X: 1, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 1, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 1, Y: 0, Z: 1, TOffsetX: 0, TOffsetY: 1},
			{X: 1, Y: 1, Z: 1, TOffsetX: 0, TOffsetY: 0},
		},
	},
	{ // South
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 0, Y: 0, Z: 1, TOffsetX: 0, TOffsetY: 1},
			{X: 0, Y: 1, Z: 1, TOffsetX: 0, TOffsetY: 0},
			{X: 1, Y: 0, Z: 1, TOffsetX: 1, TOffsetY: 1},
			{X: 1, Y: 0, Z: 1, TOffsetX: 1, TOffsetY: 1},
			{X: 0, Y: 1, Z: 1, TOffsetX: 0, TOffsetY: 0},
			{X: 1, Y: 1, Z: 1, TOffsetX: 1, TOffsetY: 0},
		},
	},
	{ // West
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 0, Y: 0, Z: 0, TOffsetX: 0, TOffsetY: 1},
			{X: 0, Y: 1, Z: 0, TOffsetX: 0, TOffsetY: 0},
			{X: 0, Y: 0, Z: 1, TOffsetX: 1, TOffsetY: 1},
			{X: 0, Y: 1, Z: 1, TOffsetX: 1, TOffsetY: 0},
		},
	},
	{ // Up
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 0, Y: 1, Z: 0, TOffsetX: 0, TOffsetY: 0},
			{X: 1, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 0, Y: 1, Z: 1, TOffsetX: 0, TOffsetY: 1},
			{X: 0, Y: 1, Z: 1, TOffsetX: 0, TOffsetY: 1},
			{X: 1, Y: 1, Z: 0, TOffsetX: 1, TOffsetY: 0},
			{X: 1, Y: 1, Z: 1, TOffsetX: 1, TOffsetY: 1},
		},
	},
	{ // Down
		indices: [6]int32{0, 1, 2, 3, 2, 1},
		verts: [6]chunkVertex{
			{X: 0, Y: 0, Z: 0, TOffsetX: 0, TOffsetY: 1},
			{X: 0, Y: 0, Z: 1, TOffsetX: 0, TOffsetY: 0},
			{X: 1, Y: 0, Z: 0, TOffsetX: 1, TOffsetY: 1},
			{X: 1, Y: 0, Z: 0, TOffsetX: 1, TOffsetY: 1},
			{X: 0, Y: 0, Z: 1, TOffsetX: 0, TOffsetY: 0},
			{X: 1, Y: 0, Z: 1, TOffsetX: 1, TOffsetY: 0},
		},
	},
}
