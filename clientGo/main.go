package main

import (
	"client/camera"
	"client/scene"
	"client/window"

	// gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	InitialiseWindow(window.SCREEN_WIDTH, window.SCREEN_HEIGHT, "CIVILISATAN", window.MAX_FPS)
	defer rl.CloseWindow()

	// Création du SceneManger
	sceneManager := scene.NewSceneManager()

	// Règle la caméra pour qu'elle soit au centre, la caméra suit le centre du plateau
	cameraOffSet := rl.Vector2{
		X: float32(window.SCREEN_WIDTH / 2),
		Y: float32(window.SCREEN_HEIGHT / 2),
	}
	camera := camera.NewCamera(cameraOffSet, sceneManager.GetGameGridCenter())

	for !rl.WindowShouldClose() {

		// Initialise la scène
		InitialiseDrawing()

		// Active la caméra 3D
		rl.BeginMode2D(*camera.Cam)
		sceneManager.DrawScene()

		// Gère le zoom de la caméra
		camera.HandlerZoom()

		rl.EndDrawing()
	}
}

func InitialiseWindow(screenWidth int32, screenHeight int32, windowName string, fps int32) {
	rl.InitWindow(screenWidth, screenHeight, windowName)
	rl.SetTargetFPS(fps)
}

func InitialiseDrawing() {
	rl.BeginDrawing()
}
