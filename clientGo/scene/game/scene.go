package game

import (
	// gui "github.com/gen2brain/raylib-go/raygui"
	"client/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	grid         *HexaGrid
	Players      []*player.Player
	ActivePlayer *player.Player
	HUD          *HUD
}

func NewGameScene() *GameScene {
	game := &GameScene{
		grid:    setupGrid(),
		Players: player.NewPlayers(2),
		HUD:     NewHUD(2),
	}
	game.ActivePlayer = game.Players[0]
	return game
}

func (g *GameScene) HandlerInput() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[0].rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer = g.Players[0]
	}

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[1].rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer = g.Players[1]
	}
}

func (g *GameScene) Draw() {
	rl.ClearBackground(rl.Black)
	g.grid.DrawGrid(g.ActivePlayer.ColorTeam)
	g.DrawHUD()
}

func (g *GameScene) GetGridCenter() rl.Vector2 {
	return g.grid.Center
}

func (g *GameScene) DrawHUD() {
	g.HUD.Draw(g.Players)
}
