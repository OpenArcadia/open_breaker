package entity

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Brick struct {
	X, Y           float32
	Width          float32
	Height         float32
	PrimaryColor   rl.Color
	SecondaryColor rl.Color
	Visible        bool
	Breakable      bool // NEW
	BreakSound     *rl.Sound
}

func NewBrick(x, y float32, breakSound *rl.Sound, breakable bool) *Brick {
	return &Brick{
		X:              x,
		Y:              y,
		Width:          70,
		Height:         30,
		PrimaryColor:   rl.DarkPurple,
		SecondaryColor: rl.Purple,
		Visible:        true,
		Breakable:      breakable,
		BreakSound:     breakSound,
	}
}

func (b *Brick) Update(ball *Ball) bool {
	if !b.Visible {
		return true
	}

	brickRect := b.GetRect()
	ballPos := rl.NewVector2(ball.X, ball.Y)

	if rl.CheckCollisionCircleRec(ballPos, ball.Radius, brickRect) {
		notShowAnimation := true
		if b.Breakable {
			rl.PlaySound(*b.BreakSound)
			b.Visible = false
			notShowAnimation = false
		}

		// Calculate the center of the ball and brick
		ballCenterX := ball.X
		ballCenterY := ball.Y
		brickCenterX := b.X + b.Width/2
		brickCenterY := b.Y + b.Height/2

		// Calculate the difference
		dx := ballCenterX - brickCenterX
		dy := ballCenterY - brickCenterY

		// Calculate overlap distances
		halfWidth := b.Width / 2
		absDx := float32(math.Abs(float64(dx)))
		absDy := float32(math.Abs(float64(dy)))

		if absDx > halfWidth && absDx > absDy {
			ball.SpeedX *= -1
		} else {
			ball.SpeedY *= -1
		}

		return notShowAnimation
	}

	return true
}

func (b *Brick) GetRect() rl.Rectangle {
	return rl.NewRectangle(b.X, b.Y, b.Width, b.Height)
}
func (b *Brick) Draw() {
	if b.Visible {
		if b.Breakable {
			rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), b.PrimaryColor)
			rl.DrawRectangle(int32(b.X+4), int32(b.Y+4), int32(b.Width-8), int32(b.Height-8), b.SecondaryColor)
		} else {
			// Grayscale for unbreakable bricks
			rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), rl.Gray)
			rl.DrawRectangle(int32(b.X+4), int32(b.Y+4), int32(b.Width-8), int32(b.Height-8), rl.LightGray)
		}
	}
}
