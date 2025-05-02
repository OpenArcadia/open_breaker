package entity

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawPlayIcon(x, y, size float32, color rl.Color) {
	// Define points of the triangle
	p1 := rl.Vector2{X: x, Y: y - size/2}
	p2 := rl.Vector2{X: x, Y: y + size/2}
	p3 := rl.Vector2{X: x + size, Y: y}

	// Draw the triangle
	rl.DrawTriangle(p1, p2, p3, color)
}
