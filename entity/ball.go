package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	X, Y   float32
	Radius float32
	SpeedX float32
	SpeedY float32
	Color  rl.Color
}

func NewBall(x, y float32) Ball {
	return Ball{
		X:      x,
		Y:      y,
		Radius: 8,
		SpeedX: 4,
		SpeedY: -4,
		Color:  rl.DarkBlue,
	}
}

func (b *Ball) Update(player Player) {
	screenWidth := rl.GetScreenWidth()

	b.X += b.SpeedX
	b.Y += b.SpeedY

	if rl.CheckCollisionCircleRec(
		rl.NewVector2(b.X, b.Y),
		b.Radius,
		player.GetRect(),
	) {
		hitPos := (b.X - (player.X + player.Width/2)) / (player.Width / 2)
		b.SpeedX = hitPos * 5
		b.SpeedY *= -1
	}

	if b.X-b.Radius <= 0 || b.X+b.Radius >= float32(screenWidth) {
		b.SpeedX *= -1
	}

	if b.Y-b.Radius <= 0 {
		b.SpeedY *= -1
	}
}

func (b Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.Color)
}

func (b Ball) Rect() rl.Rectangle {
	return rl.NewRectangle(b.X-b.Radius, b.Y-b.Radius, b.Radius*2, b.Radius*2)
}
