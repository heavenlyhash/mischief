package stuff

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

/*
	State for something that has coordinates and drifts.

	Use to compose free floating entities.  Acts under inertia plus drag;
	accelerations are impulses and zeroed after being applied.
*/
type Drifter struct {
	Position     mgl32.Vec3
	Velocity     mgl32.Vec3
	Acceleration mgl32.Vec3

	Speed float32
	Drag  float32

	Facing   mgl32.Vec3
	rotation mgl32.Vec2

	rotationSpeed float32
}

var (
	upUnit = mgl32.Vec3{0.0, 1.0, 0.0}
)

// set some default values for speed, drag, etc.
func (obj *Drifter) InitDefaults(position mgl32.Vec3) {
	obj.Position = position
	obj.Speed = 0.5
	obj.Drag = 0.5
	obj.rotationSpeed = 0.001
	obj.rotation = mgl32.Vec2{math.Pi / 1.5, math.Pi / 1.5}
	obj.Tick()
}

func (obj *Drifter) MoveForward(amount float32) {
	obj.Acceleration = obj.Acceleration.Add(obj.Facing.Mul(amount * obj.Speed))
}

func (obj *Drifter) MoveRight(amount float32) {
	obj.Acceleration = obj.Acceleration.Sub(obj.Facing.Cross(upUnit).Mul(amount * obj.Speed))
}

func (obj *Drifter) MoveUp(amount float32) {
	obj.Acceleration = obj.Acceleration.Add(upUnit.Mul(amount * obj.Speed))
}

func (obj *Drifter) Rotate(vec mgl32.Vec2) {
	obj.rotation = obj.rotation.Add(vec.Mul(obj.rotationSpeed))
}

func (obj *Drifter) Tick() {
	obj.Velocity = obj.Velocity.Add(obj.Acceleration).Mul(obj.Drag)
	obj.Position = obj.Position.Add(obj.Velocity)
	obj.Facing = obj.calculateFacingUnit()
	fmt.Printf("::  %#v\n  v %#v\n  a %#v\n  f %#v\n", obj.Position, obj.Velocity, obj.Acceleration, obj.Facing)
	obj.Acceleration = obj.Acceleration.Mul(0.0)
}

func (obj *Drifter) GetPosition() mgl32.Vec3 {
	return obj.Position
}

func (obj *Drifter) calculateFacingUnit() mgl32.Vec3 {
	// this math is using far more precision necessary...
	xsin, xcos := math.Sincos(float64(obj.rotation.X()))
	ysin, ycos := math.Sincos(float64(obj.rotation.Y()))
	return mgl32.Vec3{
		float32(ycos * xsin),
		float32(xcos),
		float32(ysin * xsin),
	}
}
