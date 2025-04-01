package hud

import "client/window"

type Action struct {
	actionButton []*RectangleButton
}

func (ac *Action) Draw() {
	for _, a := range ac.actionButton {
		a.Draw()
	}
}

func newActionHUD() *Action {
	option := []string{"Town", "Village", "Route"}
	var buttons []*RectangleButton
	for i := range 3 {
		buttons = append(
			buttons,
			newButton(
				float32(window.SCREEN_WIDTH)-buttonWidth*2*float32(1+i),
				ten,
				float32(buttonWidth),
				float32(buttonHeight),
				nil,
				option[i],
			),
		)
	}
	return &Action{
		actionButton: buttons,
	}
}
