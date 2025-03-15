package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	grid *HexaGrid
}

func NewGameScene() *GameScene {
	return &GameScene{
		grid: setupGrid(),
	}
}

func (g *GameScene) Draw() {
	rl.ClearBackground(rl.RayWhite)
	g.grid.DrawGrid()
}

func (g *GameScene) GetGridCenter() rl.Vector2 {
	return g.grid.Center
}
