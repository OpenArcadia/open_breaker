package main

import (
	"open_breaker/screens"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	rl.InitWindow(1100, 600, "brick breaker")
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
	defer rl.UnloadFont(font)

	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	screens.ChangeScreen(&screens.MenuScreen{
		Font: &font,
	})

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		screens.Update()
	}
}
