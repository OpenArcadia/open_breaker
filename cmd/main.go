package main

import (
	"open_breaker/entity"
	"open_breaker/screens"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	rl.InitWindow(800, 450, "brick breaker")
	defer rl.CloseWindow()

	_, isFlatpak := os.LookupEnv("container")

	var fontPath string
	if isFlatpak {
		fontPath = "/app/bin/assets/inter.ttf"
	} else {
		fontPath = "assets/inter.ttf"
	}

	font := rl.LoadFontEx(fontPath, 64, nil)
	game := screens.Game{
		State: screens.Menu,
		Font:  font,
	}
	rl.SetTextureFilter(game.Font.Texture, rl.FilterBilinear)

	p, ball, bricks := resetGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if game.State == screens.Menu {
			p, ball, bricks = resetGame()
			game.DrawMenu()
			continue
		}

		if game.State == screens.Over {
			p, ball, bricks = resetGame()
			game.DrawOver()
			continue
		}

		p.Update()
		ball.Update(p)

		for i := range bricks {
			bricks[i].Update(&ball)
		}

		if ball.Y >= 450 {
			game.State = screens.Over
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		p.Draw()
		ball.Draw()
		for i := range bricks {
			bricks[i].Draw()
		}

		rl.EndDrawing()
	}
}
