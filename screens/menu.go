package screens

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) DrawMenu() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	// Background color (dark gradient feel)
	rl.ClearBackground(BACKGROUND_COLOR)

	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	// Center panel
	panelWidth := int32(400)
	panelHeight := int32(400)
	panelX := screenWidth/2 - panelWidth/2
	panelY := screenHeight/2 - panelHeight/2

	panelRect := rl.NewRectangle(float32(panelX), float32(panelY), float32(panelWidth), float32(panelHeight))
	rl.DrawRectangleRounded(panelRect, 0.05, 10, rl.NewColor(10, 10, 10, 220))
	rl.DrawRectangleRoundedLines(panelRect, 0.05, 10, rl.Fade(PRIMARY_COLOR, 0.2))

	// Title text
	title := "OPEN BREAKER"
	titleSize := rl.MeasureTextEx(g.Font, title, 48, 0)
	titleX := float32(panelX) + (float32(panelWidth)-titleSize.X)/2
	rl.DrawTextEx(g.Font, title, rl.Vector2{X: titleX, Y: float32(panelY) + 30}, 48, 0, PRIMARY_COLOR)

	// Subtitle
	subtitle := "Break all the bricks to win!"
	subtitleSize := rl.MeasureTextEx(g.Font, subtitle, 20, 0)
	subtitleX := float32(panelX) + (float32(panelWidth)-subtitleSize.X)/2
	rl.DrawTextEx(g.Font, subtitle, rl.Vector2{X: subtitleX, Y: float32(panelY) + 90}, 20, 0, rl.White)

	// Buttons
	buttonWidth := int32(250)
	buttonHeight := int32(50)
	spacing := int32(20)

	firstButtonY := panelY + 150

	playBtn := rl.NewRectangle(
		float32(screenWidth/2-buttonWidth/2),
		float32(firstButtonY),
		float32(buttonWidth),
		float32(buttonHeight),
	)
	if g.DrawButton(playBtn, "Play Game", PRIMARY_COLOR, rl.Black) {
		g.State = Playing
	}

	// Other buttons (white buttons)
	labels := []string{"Settings", "Exit Game"}
	for i, label := range labels {
		btn := rl.NewRectangle(
			float32(screenWidth/2-buttonWidth/2),
			float32(firstButtonY+int32(i+1)*(buttonHeight+spacing)),
			float32(buttonWidth),
			float32(buttonHeight),
		)
		if g.DrawButton(btn, label, rl.White, rl.Black) {
			switch i {
			case 0:
				// High Scores
			case 1:
				// Settings
			case 2:
				// How to Play
			case 3:
				rl.CloseWindow()
				os.Exit(0)
			}
		}
	}

	// Footer (copyright)
	footerText := "Open Breaker | v1.0.0"
	footerSize := rl.MeasureTextEx(g.Font, footerText, 16, 0)
	rl.DrawTextEx(g.Font, footerText, rl.Vector2{
		X: float32(screenWidth/2) - footerSize.X/2,
		Y: float32(panelY + panelHeight - 30),
	}, 16, 0, rl.Fade(rl.White, 0.5))
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
