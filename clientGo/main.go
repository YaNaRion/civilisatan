package main

import (
	"client/grid"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSceen int

const (
	MENU GameSceen = iota
	GAME
)

var stateName = map[GameSceen]string{
	MENU: "Menu",
	GAME: "Game",
}

func main() {
	screenWidth := int32(1920)
	screenHeight := int32(1080)
	rl.InitWindow(screenWidth, screenHeight, "CIVILISATAN")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	sceen := MENU

	gridStr := grid.SetupGrid(rl.Vector2{
		X: float32(screenWidth) / 3,
		Y: float32(screenHeight) / 3,
	})

	log.Println(gridStr)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		switch sceen {
		case MENU:
			if rl.IsKeyPressed(rl.KeyEnter) {
				sceen = GAME
			}
		case GAME:
			gridStr.DrawGrid()
			if rl.IsKeyPressed(rl.KeyEnter) {
				sceen = MENU
			}
		}
		rl.EndDrawing()
	}
}

func DrawMenuSceen() {
	rl.DrawText("Menu, Press ENTER to enter game", 190, 200, 20, rl.LightGray)
}

func DrawGameSceen() {
	rl.DrawText(
		"THIS IS THAT GAME WINDOW, Press ENTER go back to Menu window",
		10,
		200,
		20,
		rl.LightGray,
	)
}
