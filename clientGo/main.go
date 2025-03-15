package main

import (
	"client/camera"
	"client/scene"
	"client/window"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(window.SCREEN_WIDTH, window.SCREEN_HEIGHT, "CIVILISATAN")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	// Création du SceneManger
	scene := scene.NewSceneManager(&rl.Vector2{
		X: float32(window.SCREEN_WIDTH / 2),
		Y: float32(window.SCREEN_HEIGHT / 2),
	})

	// Règle la caméra pour qu'elle soit au centre, la caméra suit le centre du plateau
	cameraOffSet := rl.Vector2{
		X: float32(window.SCREEN_WIDTH / 2),
		Y: float32(window.SCREEN_HEIGHT / 2),
	}
	camera := camera.NewCamera(cameraOffSet, scene.GetGameGridCenter())

	for !rl.WindowShouldClose() {
		// Initialise la scène
		InitialiseDrawing()

		// Active la caméra 3D
		rl.BeginMode2D(*camera.Cam)
		scene.DrawScene()

		// Gère le zoom de la caméra
		camera.HandlerZoom()

		rl.EndDrawing()
	}
}

func InitialiseDrawing() {
	rl.BeginDrawing()
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
