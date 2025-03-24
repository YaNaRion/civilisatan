package hud

import rl "github.com/gen2brain/raylib-go/raylib"

type ResetGame struct {
	Button *RectangleButton
}

func (rg *ResetGame) Draw() {
	rg.Button.Draw()
}

func newResetGameHUD() *ResetGame {
	return &ResetGame{
		Button: newButton(ten+ten*2*float32(2), ten, float32(ten), float32(ten), rl.Pink),
	}
}
