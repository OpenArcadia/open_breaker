package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameOverScreen struct {
	Font *rl.Font
}

func (g *GameOverScreen) Create() {}

func (g *GameOverScreen) Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.NewColor(10, 15, 25, 255))

	screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()

	title := "Game Over"
	subtitle1 := "Press SPACE to restart"
	subtitle2 := "Press ESC to exit"

	titleSize := 64
	subSize := 28

	// Centered horizontally
	titleWidth := rl.MeasureText(title, int32(titleSize))
	sub1Width := rl.MeasureText(subtitle1, int32(subSize))
	sub2Width := rl.MeasureText(subtitle2, int32(subSize))

	rl.DrawText(title, int32((int32(screenWidth)-titleWidth)/2), int32(screenHeight/4), int32(titleSize), rl.Red)
	rl.DrawText(subtitle1, (int32(screenWidth)-sub1Width)/2, int32(screenHeight/2), int32(subSize), rl.White)
	rl.DrawText(subtitle2, (int32(screenWidth)-sub2Width)/2, int32(screenHeight/2+50), int32(subSize), rl.White)

	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeScreen(&GameScreen{
			Font: g.Font,
		})
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		ChangeScreen(&MenuScreen{
			Font: g.Font,
		})

	}
}

func (g *GameOverScreen) Dispose() {}
