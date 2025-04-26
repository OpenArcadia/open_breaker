package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) DrawMenu() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(34, 42, 52, 255))

	buttonBg := rl.NewColor(55, 65, 81, 255)
	textColor := rl.NewColor(209, 213, 219, 255)

	buttonWidth := int32(200)
	buttonHeight := int32(50)
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	startBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2-80), float32(buttonWidth), float32(buttonHeight))
	menuBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2), float32(buttonWidth), float32(buttonHeight))
	exitBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2+80), float32(buttonWidth), float32(buttonHeight))

	if g.DrawButton(startBtn, "Start", buttonBg, textColor) {
		g.State = Playing
	}
	g.DrawButton(menuBtn, "Menu", buttonBg, textColor)
	if g.DrawButton(exitBtn, "Exit", buttonBg, textColor) {
		rl.CloseWindow()
	}
}

func (g *Game) DrawButton(rect rl.Rectangle, text string, bg rl.Color, fg rl.Color) bool {
	mousePos := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mousePos, rect)
	isClicked := isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)

	// Change color on hover
	currentBg := bg
	if isHovered {
		currentBg = rl.Fade(bg, 0.8)
	}

	// Draw button
	rl.DrawRectangleRounded(rect, 0.5, 1, currentBg)

	// Center text
	fontSize := float32(24)
	textSize := rl.MeasureTextEx(g.Font, text, fontSize, 0)
	textX := rect.X + (rect.Width-textSize.X)/2
	textY := rect.Y + (rect.Height-textSize.Y)/2

	rl.DrawTextEx(g.Font, text, rl.Vector2{X: textX, Y: textY}, fontSize, 0, fg)

	return isClicked
}
