package camera

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	Cam *rl.Camera2D
}

func NewCamera(offset rl.Vector2, target rl.Vector2) *Camera {
	return &Camera{
		Cam: &rl.Camera2D{
			Zoom:   float32(1.0),
			Target: target,
			Offset: offset,
		},
	}
}

func (c *Camera) HandlerZoom() {
	c.Cam.Zoom += rl.GetMouseWheelMove() * 0.10
	if c.Cam.Zoom < 0.1 {
		c.Cam.Zoom = 0.1
	}
}
