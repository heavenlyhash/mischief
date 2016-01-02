package controls

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"exultant.us/mischief/stuff"
)

type Camera struct {
	stuff.Drifter
}

func (cam Camera) GetLookMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		cam.Drifter.Position,
		cam.Drifter.Position.Add(cam.Drifter.Facing),
		mgl32.Vec3{0, 1, 0},
	)
}

func UpdateCameraFromKeyboard(cam *Camera, window *glfw.Window) {
	forward := window.GetKey(glfw.KeyW) == glfw.Press
	backward := window.GetKey(glfw.KeyS) == glfw.Press
	left := window.GetKey(glfw.KeyA) == glfw.Press
	right := window.GetKey(glfw.KeyD) == glfw.Press
	up := window.GetKey(glfw.KeySpace) == glfw.Press
	down := window.GetKey(glfw.KeyLeftShift) == glfw.Press

	if forward && !backward {
		cam.MoveForward(1.0)
	} else if backward && !forward {
		cam.MoveForward(-1.0)
	}

	if right && !left {
		cam.MoveRight(-1.0)
	} else if left && !right {
		cam.MoveRight(1.0)
	}

	if up && !down {
		cam.MoveUp(1.0)
	} else if down && !up {
		cam.MoveUp(-1.0)
	}
}
