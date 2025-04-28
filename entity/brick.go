package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Brick struct {
	X, Y           float32
	Width          float32
	Height         float32
	PrimaryColor   rl.Color
	SecondaryColor rl.Color
	Visible        bool
	BreakSound     *rl.Sound
}

func NewBrick(x, y float32, breakSound *rl.Sound) Brick {
	return Brick{
		X:              x,
		Y:              y,
		Width:          70,
		Height:         30,
		PrimaryColor:   rl.DarkPurple,
		SecondaryColor: rl.Purple,
		Visible:        true,
		BreakSound:     breakSound,
	}
}

func (b *Brick) Update(ball *Ball) {
	if b.Visible && rl.CheckCollisionCircleRec(
		rl.NewVector2(ball.X, ball.Y),
		ball.Radius,
		b.GetRect(),
	) {
		rl.PlaySound(*b.BreakSound)
		b.Visible = false
		ball.SpeedY *= -1
	}
}

func (b Brick) GetRect() rl.Rectangle {
	return rl.NewRectangle(b.X, b.Y, b.Width, b.Height)
}

func (b Brick) Draw() {
	if b.Visible {
		rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), b.PrimaryColor)
		rl.DrawRectangle(int32(b.X+4), int32(b.Y+4), int32(b.Width-8), int32(b.Height-8), b.SecondaryColor)
	}
}
