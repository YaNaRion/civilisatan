package hud

import "client/player"

type SwitchPlayer struct {
	Buttons []*RectangleButton
}

func newSwitchPlayerHUD(playerAmount int) *SwitchPlayer {
	return &SwitchPlayer{
		Buttons: newSwitchPlayerButton(playerAmount),
	}
}

func (sh *SwitchPlayer) Draw(players []*player.Player) {
	for _, button := range sh.Buttons {
		button.Draw()
	}
}
