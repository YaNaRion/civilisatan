package scene

import (
	"client/scene/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneManager struct {
	sceneState Scene
	gameScene  game.GameScene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		sceneState: GAME,
		gameScene:  *game.NewGameScene(),
	}
}

func (s *SceneManager) DrawScene() {
	if s.sceneState == MENU {
		s.drawGame()
	} else if s.sceneState == GAME {
		s.gameScene.Draw()
	}
}

func (s *SceneManager) drawGame() {
	rl.ClearBackground(rl.RayWhite)
	if rl.IsKeyPressed(rl.KeyEnter) {
		s.sceneState = MENU
	}

	s.gameScene.Draw()
}

func (s *SceneManager) GetGameGridCenter() rl.Vector2 {
	return s.gameScene.GetGridCenter()
}

func (s *SceneManager) HandlerInput() {
	s.gameScene.HandlerInput()
}
