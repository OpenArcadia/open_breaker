package entity

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawStar draws a simple 5-point star centered at (cx, cy)
func DrawStar(cx, cy, radius float32, color rl.Color) {
	var vertices [10]rl.Vector2

	// Two radii: outer and inner
	innerRadius := radius * 0.5
	angle := -math.Pi / 2 // Start at top

	for i := 0; i < 10; i++ {
		r := radius
		if i%2 != 0 {
			r = innerRadius
		}
		vertices[i] = rl.NewVector2(
			cx+float32(math.Cos(angle))*r,
			cy+float32(math.Sin(angle))*r,
		)
		angle += math.Pi / 5 // 36 degrees
	}

	// Draw filled star using triangle fan
	for i := 1; i < 9; i++ {
		rl.DrawTriangle(vertices[0], vertices[i], vertices[i+1], color)
	}
}
