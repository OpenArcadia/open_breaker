package screens

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FinishScreen struct {
	Font         *rl.Font
	StartTime    time.Time
	EndTime      time.Time
	displayTime  string
	CurrentLevel LevelName
	NextLevel    LevelName
}

func (f *FinishScreen) Create() {
	duration := f.EndTime.Sub(f.StartTime)
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	f.displayTime = fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func (f *FinishScreen) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	rl.ClearBackground(BACKGROUND_COLOR)

	cardWidth := float32(400)
	cardHeight := float32(460)
	cardX := float32(screenWidth)/2 - cardWidth/2
	cardY := float32(screenHeight)/2 - cardHeight/2

	rl.DrawRectangleRounded(rl.NewRectangle(cardX, cardY, cardWidth, cardHeight), 0.05, 10, rl.NewColor(20, 24, 40, 255))

	congratsText := "Congratulations!"
	congratsSize := rl.MeasureTextEx(*f.Font, congratsText, 40, 0)
	rl.DrawTextEx(*f.Font, congratsText, rl.Vector2{
		X: cardX + cardWidth/2 - congratsSize.X/2,
		Y: cardY + 40,
	}, 40, 0, PRIMARY_COLOR)

	subText := "You finished the game!"
	subSize := rl.MeasureTextEx(*f.Font, subText, 20, 0)
	rl.DrawTextEx(*f.Font, subText, rl.Vector2{
		X: cardX + cardWidth/2 - subSize.X/2,
		Y: cardY + 100,
	}, 20, 0, rl.LightGray)

	scoreLabel := "Score"
	timeLabel := "Time"
	scoreValue := "8000" // Replace with actual score

	labelFontSize := float32(16)
	valueFontSize := float32(24)

	boxWidth := float32(140)
	boxHeight := float32(70)
	boxSpacing := float32(20)

	scoreBox := rl.NewRectangle(cardX+boxSpacing, cardY+150, boxWidth, boxHeight)
	timeBox := rl.NewRectangle(cardX+cardWidth-boxSpacing-boxWidth, cardY+150, boxWidth, boxHeight)

	rl.DrawRectangleRounded(scoreBox, 0.1, 5, rl.NewColor(10, 12, 20, 255))
	rl.DrawRectangleRounded(timeBox, 0.1, 5, rl.NewColor(10, 12, 20, 255))

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

	timeLabelSize := rl.MeasureTextEx(*f.Font, timeLabel, labelFontSize, 0)
	rl.DrawTextEx(*f.Font, timeLabel, rl.Vector2{
		X: timeBox.X + timeBox.Width/2 - timeLabelSize.X/2,
		Y: timeBox.Y + 5,
	}, labelFontSize, 0, rl.LightGray)

	timeValueSize := rl.MeasureTextEx(*f.Font, f.displayTime, valueFontSize, 0)
	rl.DrawTextEx(*f.Font, f.displayTime, rl.Vector2{
		X: timeBox.X + timeBox.Width/2 - timeValueSize.X/2,
		Y: timeBox.Y + 30,
	}, valueFontSize, 0, rl.NewColor(255, 193, 7, 255))

	// Buttons
	buttonWidth := cardWidth - boxSpacing*2
	buttonHeight := float32(50)
	buttonGap := float32(15)

	restartButton := rl.NewRectangle(cardX+boxSpacing, cardY+250, buttonWidth, buttonHeight)
	nextStageButton := rl.NewRectangle(cardX+boxSpacing, cardY+250+buttonHeight+buttonGap, buttonWidth, buttonHeight)
	exitButton := rl.NewRectangle(cardX+boxSpacing, cardY+250+2*(buttonHeight+buttonGap), buttonWidth, buttonHeight)

	if drawButton(f.Font, restartButton, "Restart (R)", 20, PRIMARY_COLOR, rl.Fade(PRIMARY_COLOR, 0.8)) {
		ChangeScreen(&GameScreen{Font: f.Font, CurrentLevel: f.CurrentLevel})
	}

	if f.NextLevel != "" && drawButton(f.Font, nextStageButton, "Next Stage", 20, PRIMARY_COLOR, rl.Fade(PRIMARY_COLOR, 0.8)) {
		ChangeScreen(&GameScreen{Font: f.Font, CurrentLevel: f.NextLevel})
	}

	if drawButton(f.Font, exitButton, "Exit (ESC)", 20, PRIMARY_COLOR, rl.Fade(PRIMARY_COLOR, 0.8)) {
		ChangeScreen(&MenuScreen{Font: f.Font})
	}

	// Optional: support keyboard shortcuts
	if rl.IsKeyPressed(rl.KeyR) {
		ChangeScreen(&GameScreen{Font: f.Font})
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		ChangeScreen(&MenuScreen{Font: f.Font})
	}
}

func (f *FinishScreen) Dispose() {}

func drawButton(font *rl.Font, rect rl.Rectangle, text string, fontSize float32, normalColor, hoverColor rl.Color) bool {
	mousePos := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mousePos, rect)
	color := normalColor
	if isHovered {
		color = hoverColor
	}
	rl.DrawRectangleRounded(rect, 0.2, 5, color)

	textSize := rl.MeasureTextEx(*font, text, fontSize, 0)
	rl.DrawTextEx(*font, text, rl.Vector2{
		X: rect.X + rect.Width/2 - textSize.X/2,
		Y: rect.Y + rect.Height/2 - textSize.Y/2,
	}, fontSize, 0, rl.Black)

	return isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)
}
