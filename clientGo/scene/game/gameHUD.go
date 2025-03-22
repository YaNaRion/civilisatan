package game

import (
	"client/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const ten float32 = 30

type HUD struct {
	SwitchPlayer *SwitchPlayer
	ResetGame    *ResetGame
}

type SwitchPlayer struct {
	Buttons []*RectangleButton
}

type ResetGame struct {
	Button *RectangleButton
}

type RectangleButton struct {
	rec   *rl.Rectangle
	color rl.Color
}

func NewHUD(playerAmount int) *HUD {
	return &HUD{
		SwitchPlayer: newSwitchPlayerHUD(playerAmount),
		ResetGame:    newResetGameHUD(),
	}
}

func (h *HUD) Draw(players []*player.Player) {
	h.SwitchPlayer.Draw(players)
	h.ResetGame.Draw()
}

func newSwitchPlayerHUD(playerAmount int) *SwitchPlayer {
	return &SwitchPlayer{
		Buttons: newSwitchPlayerButton(playerAmount),
	}
}

func newResetGameHUD() *ResetGame {
	return &ResetGame{
		Button: newButton(ten+ten*2*float32(2), ten, float32(ten), float32(ten), rl.Pink),
	}
}

func (rg *ResetGame) Draw() {
	rg.Button.Draw()
}

func (sh *SwitchPlayer) Draw(players []*player.Player) {
	for _, button := range sh.Buttons {
		button.Draw()
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

func (b *RectangleButton) Draw() {
	rl.DrawRectangle(
		b.rec.ToInt32().X,
		b.rec.ToInt32().Y,
		b.rec.ToInt32().Width,
		b.rec.ToInt32().Height,
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
		rec:   &rec,
		color: color,
	}
}
