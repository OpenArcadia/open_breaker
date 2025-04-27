package main

import (
	"open_breaker/entity"
	"open_breaker/screens"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func resetGame(breakSound, bounceSound *rl.Sound) (entity.Player, entity.Ball, []entity.Brick) {
	p := entity.NewPlayer(350, 400)
	ball := entity.NewBall(400, 300, bounceSound)

	bricks := []entity.Brick{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			bricks = append(bricks, entity.NewBrick(70+float32(i)*70, 50+float32(j)*30, breakSound))
		}
	}

	return p, ball, bricks
}

func Filter[T any](input []T, test func(T) bool) []T {
	var result []T
	for _, v := range input {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	rl.InitWindow(800, 450, "brick breaker")
	rl.InitAudioDevice()
	rl.SetExitKey(0)
	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	_, isFlatpak := os.LookupEnv("container")

	var basePath string
	if isFlatpak {
		basePath = "/app/bin/assets/"
	} else {
		basePath = "assets/"
	}

	font := rl.LoadFontEx(basePath+"inter.ttf", 64, nil)

	breakSound := rl.LoadSound(basePath + "music/break.ogg")
	rl.SetSoundVolume(breakSound, 1.0)
	bounceSound := rl.LoadSound(basePath + "music/bounce.ogg")
	rl.SetSoundVolume(bounceSound, 1.0)
	defer rl.UnloadSound(breakSound)
	defer rl.UnloadSound(bounceSound)
	defer rl.UnloadFont(font)

	game := screens.Game{
		State: screens.Menu,
		Font:  font,
	}
	rl.SetTextureFilter(game.Font.Texture, rl.FilterBilinear)

	p, ball, bricks := resetGame(&breakSound, &bounceSound)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if game.State == screens.Finished {
			p, ball, bricks = resetGame(&breakSound, &bounceSound)
			game.DrawFinishScreen()
			continue
		}

		if game.State == screens.Menu {
			p, ball, bricks = resetGame(&breakSound, &bounceSound)
			game.DrawMenu()
			continue
		}

		if game.State == screens.Over {
			p, ball, bricks = resetGame(&breakSound, &bounceSound)
			game.DrawOver()
			continue
		}

		p.Update()
		ball.Update(p)

		for i := range bricks {
			bricks[i].Update(&ball)
		}

		bricks = Filter(bricks, func(b entity.Brick) bool {
			return b.Visible
		})

		if len(bricks) == 0 {
			game.State = screens.Finished
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
