package game

import (
	"client/player"
	"client/scene/game/hud"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	grid             *HexaGrid
	Players          []*player.Player
	ActivePlayer     *player.Player
	CurrentDicevalue int
	HUD              *hud.HUD
}

func NewGameScene() *GameScene {
	gameState := "DEBUT DE LA GAME"
	game := &GameScene{
		grid:    setupGrid(),
		Players: player.NewPlayers(2),
		HUD:     hud.NewHUD(2, &gameState),
	}
	game.ActivePlayer = game.Players[0]
	return game
}

func (g *GameScene) HandlerInput() {
	// Switch to player0
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[0].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer = g.Players[0]
	}

	// Switch to player1
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.SwitchPlayer.Buttons[1].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer = g.Players[1]
	}

	// ResetGame
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.ResetGame.Button.Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ResetGame()
	}

	// Town button
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.Action.ActionButton[0].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer.Action = player.UpgradeTownCenter
	}

	// Village button
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.Action.ActionButton[1].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer.Action = player.PlaceVillage
	}

	// Route button
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.Action.ActionButton[2].Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.ActivePlayer.Action = player.PlaceRoute
	}

	// Dice button
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), *g.HUD.Dice.Button.Rec) &&
		rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		g.CurrentDicevalue = g.HUD.Dice.RolDice()
	}
}

func (g *GameScene) Draw() {
	rl.ClearBackground(rl.Black)
	g.grid.DrawGrid(g.ActivePlayer)
	g.DrawHUD()
}

func (g *GameScene) GetGridCenter() rl.Vector2 {
	return g.grid.Center
}

func (g *GameScene) DrawHUD() {
	g.HUD.Draw(g.Players)
}

func (g *GameScene) ResetGame() {
	g.ActivePlayer = g.Players[0]
	g.grid.ResetGrid()
	g.grid.PopulateGrid()
}
