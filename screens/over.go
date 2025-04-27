package screens

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) DrawOver() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Blank)

	rl.DrawText("Game Over", 300, 100, 42, rl.Red)
	rl.DrawText("Press space to restart", 250, 200, 28, rl.Red)
	rl.DrawText("Press escape to exit", 270, 250, 28, rl.Red)
	if rl.IsKeyPressed(rl.KeySpace) {
		g.State = Playing
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		g.State = Menu
	}
}
