package main

import (
	"open_breaker/entity"
	"open_breaker/screens"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func removeAtIndex(slice []entity.Brick, index int) []entity.Brick {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func resetGame() (entity.Player, entity.Ball, []entity.Brick) {
	p := entity.NewPlayer(350, 400)
	ball := entity.NewBall(400, 300)

	bricks := []entity.Brick{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			bricks = append(bricks, entity.NewBrick(70+float32(i)*70, 50+float32(j)*30))
		}
	}

	return p, ball, bricks
}

func main() {
	gameOver := false
	rl.InitWindow(800, 450, "brick breaker")
	defer rl.CloseWindow()
	font := rl.LoadFontEx("assets/inter.ttf", 64, nil)
	game := screens.Game{
		State: screens.Menu,
		Font:  font,
	}
	rl.SetTextureFilter(game.Font.Texture, rl.FilterBilinear)

	p, ball, bricks := resetGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if game.State == screens.Menu {
			game.DrawMenu()
			continue
		}

		if gameOver {
			if rl.IsKeyPressed(rl.KeySpace) {
				p, ball, bricks = resetGame()
				gameOver = false
			}
		} else {
			p.Update()
			ball.Update(p)

			for i := range bricks {
				bricks[i].Update(&ball)
			}

			if ball.Y >= 450 {
				gameOver = true
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		if gameOver {
			rl.DrawText("GAME OVER", 280, 180, 40, rl.Red)
			rl.DrawText("Press SPACE to Restart", 270, 240, 20, rl.DarkGray)
		} else {
			p.Draw()
			ball.Draw()
			for i := range bricks {
				bricks[i].Draw()
			}
		}

		rl.EndDrawing()
	}
}
