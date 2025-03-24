package hud

import rl "github.com/gen2brain/raylib-go/raylib"

type RectangleButton struct {
	Rec   *rl.Rectangle
	color rl.Color
}

func (b *RectangleButton) Draw() {
	rl.DrawRectangle(
		b.Rec.ToInt32().X,
		b.Rec.ToInt32().Y,
		b.Rec.ToInt32().Width,
		b.Rec.ToInt32().Height,
		b.color,
	)
}

func newButton(
	x float32,
	y float32,
	width float32,
	height float32,
	color rl.Color,
) *RectangleButton {
	rec := rl.NewRectangle(
		x,
		y,
		width,
		height,
	)
	return &RectangleButton{
		Rec:   &rec,
		color: color,
	}
}

func newSwitchPlayerButton(playerAmount int) []*RectangleButton {
	var col []rl.Color
	col = append(col, rl.Green)
	col = append(col, rl.Red)

	var buttons []*RectangleButton
	for i := range playerAmount {
		buttons = append(
			buttons,
			newButton(ten+ten*2*float32(i), ten, float32(ten), float32(ten), col[i]),
		)
	}
	return buttons
}
