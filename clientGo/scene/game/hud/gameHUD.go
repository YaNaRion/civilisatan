package hud

import (
	"client/player"
)

const ten float32 = 30

const (
	buttonWidth  float32 = 60
	buttonHeight float32 = 30
)

type HUD struct {
	SwitchPlayer *SwitchPlayer
	ResetGame    *ResetGame
	Action       *Action
}

func NewHUD(playerAmount int) *HUD {
	return &HUD{
		SwitchPlayer: newSwitchPlayerHUD(playerAmount),
		ResetGame:    newResetGameHUD(),
		Action:       newActionHUD(),
	}
}

func (h *HUD) Draw(players []*player.Player) {
	h.SwitchPlayer.Draw(players)
	h.ResetGame.Draw()
	h.Action.Draw()
}
