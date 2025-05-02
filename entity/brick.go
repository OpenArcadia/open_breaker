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

func NewBrick(width, height float32, x, y float32, breakSound *rl.Sound, breakable bool) *Brick {
	return &Brick{
		X:              x,
		Y:              y,
		Width:          width,
		Height:         height,
		PrimaryColor:   rl.DarkPurple,
		SecondaryColor: rl.Purple,
		Visible:        true,
		Breakable:      breakable,
		BreakSound:     breakSound,
	}
}

func (b *Brick) Update(ball *Ball) bool {
	brickRect := b.GetRect()
	ballPos := rl.NewVector2(ball.X, ball.Y)

	if rl.CheckCollisionCircleRec(ballPos, ball.Radius, brickRect) {
		noAnimation := true
		if b.Breakable {
			rl.PlaySound(*b.BreakSound)
			b.Visible = false
			noAnimation = false
		}

		// Find centers
		ballCenterX := ball.X
		ballCenterY := ball.Y
		brickCenterX := b.X + b.Width/2
		brickCenterY := b.Y + b.Height/2

		// Deltas
		dx := ballCenterX - brickCenterX
		dy := ballCenterY - brickCenterY
		absDx := float32(math.Abs(float64(dx)))
		absDy := float32(math.Abs(float64(dy)))

		// Half-dimensions
		halfW := b.Width / 2
		halfH := b.Height / 2

		// Calculate overlap along X and Y
		overlapX := halfW + ball.Radius - absDx
		overlapY := halfH + ball.Radius - absDy

		if overlapX < overlapY {
			// Push out horizontally
			if dx > 0 {
				ball.X += overlapX
			} else {
				ball.X -= overlapX
			}
			ball.SpeedX *= -1
		} else {
			// Push out vertically
			if dy > 0 {
				ball.Y += overlapY
			} else {
				ball.Y -= overlapY
			}
			ball.SpeedY *= -1
		}

		return noAnimation
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
