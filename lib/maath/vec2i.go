package maath

import "github.com/go-gl/mathgl/mgl32"

type Vec2i [2]int

// Elem extracts the elements of the vector for direct value assignment.
func (v Vec2i) Elem() (x, y int) {
	return v[0], v[1]
}

func (v Vec2i) X() int { return v[0] }
func (v Vec2i) Y() int { return v[1] }

func (v1 Vec2i) Add(v2 Vec2i) Vec2i { return Vec2i{v1[0] + v2[0], v1[1] + v2[1]} }
func (v1 Vec2i) Sub(v2 Vec2i) Vec2i { return Vec2i{v1[0] - v2[0], v1[1] - v2[1]} }
func (v1 Vec2i) Mul(val int) Vec2i  { return Vec2i{v1[0] * val, v1[1] * val} }
func (v1 Vec2i) Div(val int) Vec2i  { return Vec2i{v1[0] / val, v1[1] / val} }
func (v1 Vec2i) Mod(val int) Vec2i  { return Vec2i{v1[0] % val, v1[1] % val} }

func (v Vec2i) Vec2() mgl32.Vec2 {
	return mgl32.Vec2{float32(v[0]), float32(v[1])}
}
