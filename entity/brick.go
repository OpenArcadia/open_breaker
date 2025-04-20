package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	X, Y    float32
	Width   float32
	Height  float32
	Color   rl.Color
	Visible bool
}

func NewBrick(x, y float32) Brick {
	return Brick{
		X:       x,
		Y:       y,
		Width:   60,
		Height:  20,
		Color:   rl.Red,
		Visible: true,
	}
}

func (b *Brick) Update(ball *Ball) {
	if b.Visible && rl.CheckCollisionCircleRec(
		rl.NewVector2(ball.X, ball.Y),
		ball.Radius,
		b.GetRect(),
	) {
		b.Visible = false
		ball.SpeedY *= -1
	}
}

func (b Brick) GetRect() rl.Rectangle {
	return rl.NewRectangle(b.X, b.Y, b.Width, b.Height)
}

func (b Brick) Draw() {
	if b.Visible {
		rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), b.Color)
	}
}
