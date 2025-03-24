package hud

import (
	"client/player"
)

const ten float32 = 30

type HUD struct {
	SwitchPlayer *SwitchPlayer
	ResetGame    *ResetGame
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
