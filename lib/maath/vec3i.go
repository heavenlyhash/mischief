package maath

import "github.com/go-gl/mathgl/mgl32"

type Vec3i [3]int

// Elem extracts the elements of the vector for direct value assignment.
func (v Vec3i) Elem() (x, y, z int) {
	return v[0], v[1], v[2]
}

func (v Vec3i) X() int { return v[0] }
func (v Vec3i) Y() int { return v[1] }
func (v Vec3i) Z() int { return v[2] }

func (v1 Vec3i) Add(v2 Vec3i) Vec3i { return Vec3i{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]} }
func (v1 Vec3i) Sub(v2 Vec3i) Vec3i { return Vec3i{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]} }
func (v1 Vec3i) Mul(val int) Vec3i  { return Vec3i{v1[0] * val, v1[1] * val, v1[2] * val} }
func (v1 Vec3i) Div(val int) Vec3i  { return Vec3i{v1[0] / val, v1[1] / val, v1[2] / val} }
func (v1 Vec3i) Mod(val int) Vec3i  { return Vec3i{v1[0] % val, v1[1] % val, v1[2] % val} }

func (v Vec3i) Vec3() mgl32.Vec3 {
	return mgl32.Vec3{float32(v[0]), float32(v[1]), float32(v[2])}
}
