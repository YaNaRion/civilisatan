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
	Dice         *Dice
	GameStatus   *GameStatus
}

func NewHUD(playerAmount int, gameStatus *string) *HUD {
	return &HUD{
		SwitchPlayer: newSwitchPlayerHUD(playerAmount),
		ResetGame:    newResetGameHUD(),
		Action:       newActionHUD(),
		Dice:         newDiceHUD(),
		GameStatus:   newGameStatusHUD(gameStatus),
	}
}

func (h *HUD) Draw(players []*player.Player) {
	h.SwitchPlayer.Draw(players)
	h.ResetGame.Draw()
	h.Action.Draw()
	h.Dice.Draw()
	h.GameStatus.Draw()
}
