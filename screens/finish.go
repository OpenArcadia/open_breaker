package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FinishScreen struct {
	Font *rl.Font
}

func (f *FinishScreen) Create() {}

func (f *FinishScreen) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	// Background: dark gradient simulation (solid color for now)
	rl.ClearBackground(BACKGROUND_COLOR) // Very dark blue

	// Center card
	cardWidth := float32(400)
	cardHeight := float32(400)
	cardX := float32(screenWidth)/2 - cardWidth/2
	cardY := float32(screenHeight)/2 - cardHeight/2

	rl.DrawRectangleRounded(rl.NewRectangle(cardX, cardY, cardWidth, cardHeight), 0.05, 10, rl.NewColor(20, 24, 40, 255))

	// Trophy emoji (optional: if you want you could draw an icon here)
	// rl.DrawTextEx(*f.Font, "🏆", rl.Vector2{X: cardX + cardWidth/2 - 30, Y: cardY + 20}, 60, 0, rl.Gold)

	// Congratulations
	congratsText := "Congratulations!"
	congratsSize := rl.MeasureTextEx(*f.Font, congratsText, 40, 0)
	rl.DrawTextEx(*f.Font, congratsText, rl.Vector2{
		X: cardX + cardWidth/2 - congratsSize.X/2,
		Y: cardY + 40,
	}, 40, 0, PRIMARY_COLOR) // Gold color

	// Subtext
	subText := "You finished the game!"
	subSize := rl.MeasureTextEx(*f.Font, subText, 20, 0)
	rl.DrawTextEx(*f.Font, subText, rl.Vector2{
		X: cardX + cardWidth/2 - subSize.X/2,
		Y: cardY + 100,
	}, 20, 0, rl.LightGray)

	// Score and Time
	scoreLabel := "Score"
	timeLabel := "Time"
	scoreValue := "8000" // Replace with actual score
	timeValue := "02:45" // Replace with actual time

	labelFontSize := float32(16)
	valueFontSize := float32(24)

	// Score box
	boxWidth := float32(140)
	boxHeight := float32(70)
	boxSpacing := float32(20)

	scoreBox := rl.NewRectangle(cardX+boxSpacing, cardY+150, boxWidth, boxHeight)
	timeBox := rl.NewRectangle(cardX+cardWidth-boxSpacing-boxWidth, cardY+150, boxWidth, boxHeight)

	rl.DrawRectangleRounded(scoreBox, 0.1, 5, rl.NewColor(10, 12, 20, 255))
	rl.DrawRectangleRounded(timeBox, 0.1, 5, rl.NewColor(10, 12, 20, 255))

	// Draw score text
	scoreLabelSize := rl.MeasureTextEx(*f.Font, scoreLabel, labelFontSize, 0)
	rl.DrawTextEx(*f.Font, scoreLabel, rl.Vector2{
		X: scoreBox.X + scoreBox.Width/2 - scoreLabelSize.X/2,
		Y: scoreBox.Y + 5,
	}, labelFontSize, 0, rl.LightGray)

	scoreValueSize := rl.MeasureTextEx(*f.Font, scoreValue, valueFontSize, 0)
	rl.DrawTextEx(*f.Font, scoreValue, rl.Vector2{
		X: scoreBox.X + scoreBox.Width/2 - scoreValueSize.X/2,
		Y: scoreBox.Y + 30,
	}, valueFontSize, 0, rl.NewColor(255, 193, 7, 255))

	// Draw time text
	timeLabelSize := rl.MeasureTextEx(*f.Font, timeLabel, labelFontSize, 0)
	rl.DrawTextEx(*f.Font, timeLabel, rl.Vector2{
		X: timeBox.X + timeBox.Width/2 - timeLabelSize.X/2,
		Y: timeBox.Y + 5,
	}, labelFontSize, 0, rl.LightGray)

	timeValueSize := rl.MeasureTextEx(*f.Font, timeValue, valueFontSize, 0)
	rl.DrawTextEx(*f.Font, timeValue, rl.Vector2{
		X: timeBox.X + timeBox.Width/2 - timeValueSize.X/2,
		Y: timeBox.Y + 30,
	}, valueFontSize, 0, rl.NewColor(255, 193, 7, 255))

	// Buttons
	buttonWidth := cardWidth - boxSpacing*2
	buttonHeight := float32(50)

	restartButton := rl.NewRectangle(cardX+boxSpacing, cardY+250, buttonWidth, buttonHeight)
	exitButton := rl.NewRectangle(cardX+boxSpacing, cardY+310, buttonWidth, buttonHeight)

	// Restart button
	rl.DrawRectangleRounded(restartButton, 0.2, 5, PRIMARY_COLOR)
	restartText := "Restart (R)"
	restartSize := rl.MeasureTextEx(*f.Font, restartText, 20, 0)
	rl.DrawTextEx(*f.Font, restartText, rl.Vector2{
		X: restartButton.X + restartButton.Width/2 - restartSize.X/2,
		Y: restartButton.Y + restartButton.Height/2 - restartSize.Y/2,
	}, 20, 0, rl.Black)

	// Exit button
	rl.DrawRectangleRounded(exitButton, 0.2, 5, PRIMARY_COLOR)
	exitText := "Exit (ESC)"
	exitSize := rl.MeasureTextEx(*f.Font, exitText, 20, 0)
	rl.DrawTextEx(*f.Font, exitText, rl.Vector2{
		X: exitButton.X + exitButton.Width/2 - exitSize.X/2,
		Y: exitButton.Y + exitButton.Height/2 - exitSize.Y/2,
	}, 20, 0, rl.LightGray)

	// Button click handling
	if rl.IsKeyPressed(rl.KeyR) {
		ChangeScreen(&GameScreen{
			Font: f.Font,
		})
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		ChangeScreen(&MenuScreen{
			Font: f.Font,
		})
	}
}

func (f *FinishScreen) Dispose() {}
