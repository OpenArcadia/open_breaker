package screens

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) DrawMenu() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	// Darker background color
	rl.ClearBackground(rl.NewColor(25, 25, 30, 255))

	// Draw the title centered on the screen
	rl.DrawTextEx(g.Font, "Brick Breaker", rl.Vector2{X: 250, Y: 100}, 64, 0, rl.NewColor(209, 213, 219, 255))

	// Button styles
	buttonBg := rl.NewColor(55, 65, 81, 255)
	textColor := rl.NewColor(209, 213, 219, 255)

	// Adjust button size and positions
	buttonWidth := int32(200) // Increased button width
	buttonHeight := int32(40) // Increased button height
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	// Button positions centered relative to the title
	startBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2+30), float32(buttonWidth), float32(buttonHeight))
	menuBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2+90), float32(buttonWidth), float32(buttonHeight))
	exitBtn := rl.NewRectangle(float32(screenWidth/2-buttonWidth/2), float32(screenHeight/2+150), float32(buttonWidth), float32(buttonHeight))

	// Draw buttons and detect clicks
	if g.DrawButton(startBtn, "Start", buttonBg, textColor) {
		g.State = Playing
	}
	if g.DrawButton(menuBtn, "Settings", buttonBg, textColor) {
		// Open settings menu
	}
	if g.DrawButton(exitBtn, "Exit", buttonBg, textColor) {
		rl.CloseWindow() // Exit the game
		os.Exit(0)
	}
}
func (g *Game) DrawButton(rect rl.Rectangle, text string, bg rl.Color, fg rl.Color) bool {
	mousePos := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mousePos, rect)
	isClicked := isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)

	// Hover effect: Lighten color and add shadow
	currentBg := bg
	if isHovered {
		currentBg = rl.Fade(bg, 0.8)
	}

	// Button shadow for depth
	shadowOffset := float32(4)
	shadowColor := rl.NewColor(0, 0, 0, 150)

	// Draw shadow first, under the button
	if isHovered || isClicked {
		rl.DrawRectangleRounded(rl.NewRectangle(rect.X+shadowOffset, rect.Y+shadowOffset, rect.Width, rect.Height), 0.5, 1, shadowColor)
	}

	// Draw the main button with rounded corners
	rl.DrawRectangleRounded(rect, 0.5, 1, currentBg)

	// Center the text on the button
	fontSize := float32(28)
	textSize := rl.MeasureTextEx(g.Font, text, fontSize, 0)
	textX := rect.X + (rect.Width-textSize.X)/2
	textY := rect.Y + (rect.Height-textSize.Y)/2

	// Draw text on the button
	rl.DrawTextEx(g.Font, text, rl.Vector2{X: textX, Y: textY}, fontSize, 0, fg)

	return isClicked
}
