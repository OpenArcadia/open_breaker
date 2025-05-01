package screens

import (
	"open_breaker/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelName string

const (
	UNKNOWN_LEVEL LevelName = ""
	LEVEL_ONE     LevelName = "Level 1"
	LEVEL_TWO     LevelName = "Level 2"
	LEVEL_THREE   LevelName = "Level 3"
	LEVEL_FOUR    LevelName = "Level 4"
	LEVEL_FIVE    LevelName = "Level 5"
	LEVEL_SIX     LevelName = "Level 6"
	LEVEL_SEVEN   LevelName = "Level 7"
	LEVEL_EIGHT   LevelName = "Level 8"
	LEVEL_NINE    LevelName = "Level 9"
	LEVEL_TEN     LevelName = "Level 10"
)

func (ln *LevelName) ToString() string {
	if ln == nil {
		return ""
	}
	return string(*ln)
}

type Level struct {
	Name     LevelName
	Unlocked bool
	Stars    int // 0–3
}

type LevelScreen struct {
	Font *rl.Font
}

func (g *LevelScreen) Create() {}

func (g *LevelScreen) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(10, 10, 10, 255))

	drawBackToMenu(g)

	levels := []Level{
		{"Level 1", true, 3},
		{"Level 2", true, 2},
		{"Level 3", true, 1},
		{"Level 4", true, 0},
		{"Level 5", true, 0},
		{"Level 6", false, 0},
		{"Level 7", false, 0},
		{"Level 8", false, 0},
	}

	drawLevelCards(g, levels)
}

func (g *LevelScreen) Dispose() {}

func drawBackToMenu(g *LevelScreen) {
	const fontSize = 24
	text := "<  Back to Menu"
	pos := rl.Vector2{X: 20, Y: 30}
	textSize := rl.MeasureTextEx(*g.Font, text, fontSize, 0)

	rl.DrawTextEx(*g.Font, text, pos, fontSize, 0, rl.Gray)

	backRect := rl.Rectangle{
		X: pos.X, Y: pos.Y,
		Width: textSize.X, Height: textSize.Y,
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&
		rl.CheckCollisionPointRec(rl.GetMousePosition(), backRect) {
		ChangeScreen(&MenuScreen{
			Font: g.Font,
		})
	}
}

func drawLevelCards(g *LevelScreen, levels []Level) {
	screenWidth := float32(rl.GetScreenWidth())

	const cardWidth = 600
	const cardHeight = 60
	const startX = float32(80)
	const startY = float32(100)
	const spacing = 20

	for i, level := range levels {
		y := startY + float32(i)*(cardHeight+spacing)
		rect := rl.Rectangle{X: screenWidth - (startX + cardWidth), Y: y, Width: cardWidth, Height: cardHeight}

		drawLevelCard(g, level, rect)
	}
}

func drawLevelCard(g *LevelScreen, level Level, rect rl.Rectangle) {
	// Background
	bgColor := rl.NewColor(30, 30, 30, 255)
	if !level.Unlocked {
		bgColor = rl.NewColor(20, 20, 20, 255)
	}
	rl.DrawRectangleRounded(rect, 0.1, 10, bgColor)

	// Level text
	textColor := rl.White
	if !level.Unlocked {
		textColor = rl.DarkGray
	}
	rl.DrawTextEx(*g.Font, level.Name.ToString(),
		rl.Vector2{X: rect.X + 20, Y: rect.Y + 16}, 24, 0, textColor)

	// Stars
	for s := 0; s < 3; s++ {
		starX := rect.X + 220 + float32(s)*24
		starColor := rl.Gray
		if level.Stars > s {
			starColor = rl.Yellow
		}
		entity.DrawStar(starX, rect.Y+35, 10, starColor)
	}

	// Play button
	playBtn := rl.Rectangle{
		X:     rect.X + rect.Width - 50,
		Y:     rect.Y + 15,
		Width: 30, Height: 30,
	}
	playColor := rl.NewColor(200, 160, 255, 255)
	if !level.Unlocked {
		playColor = rl.Fade(playColor, 0.3)
	}
	rl.DrawRectangleRounded(playBtn, 0.3, 10, playColor)
	rl.DrawTextEx(*g.Font, "▶",
		rl.Vector2{X: playBtn.X + 6, Y: playBtn.Y + 2}, 20, 0, rl.Black)

	// Play button interaction (optional)
	if level.Unlocked &&
		rl.CheckCollisionPointRec(rl.GetMousePosition(), playBtn) &&
		rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ChangeScreen(&GameScreen{
			CurrentLevel: level.Name,
			Font:         g.Font,
		})
	}
}
