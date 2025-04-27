package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	X, Y        float32
	Radius      float32
	SpeedX      float32
	SpeedY      float32
	Color       rl.Color
	BounceSound *rl.Sound
}

type BallTrail struct {
	Position rl.Vector2
	Alpha    uint8
}

var ballTrail []BallTrail

const MaxTrailLength = 10

func NewBall(x, y float32, bounceSound *rl.Sound) Ball {
	return Ball{
		X:           x,
		Y:           y,
		Radius:      8,
		SpeedX:      4,
		SpeedY:      -4,
		Color:       rl.Beige,
		BounceSound: bounceSound,
	}
}

func (b *Ball) Update(player Player) {
	if len(ballTrail) > MaxTrailLength {
		ballTrail = ballTrail[1:]
	}

	ballTrail = append(ballTrail, BallTrail{
		Position: rl.Vector2{X: b.X, Y: b.Y},
		Alpha:    255,
	})

	screenWidth := rl.GetScreenWidth()

	b.X += b.SpeedX
	b.Y += b.SpeedY

	if rl.CheckCollisionCircleRec(
		rl.NewVector2(b.X, b.Y),
		b.Radius,
		player.GetRect(),
	) {
		hitPos := (b.X - (player.X + player.Width/2)) / (player.Width / 2)
		rl.PlaySound(*b.BounceSound)
		b.SpeedX = hitPos * 5
		b.SpeedY *= -1
	}

	if b.X-b.Radius <= 0 || b.X+b.Radius >= float32(screenWidth) {
		rl.PlaySound(*b.BounceSound)
		b.SpeedX *= -1
	}

	if b.Y-b.Radius <= 0 {
		rl.PlaySound(*b.BounceSound)
		b.SpeedY *= -1
	}
}

func (b Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.Color)

	for i, trail := range ballTrail {
		fadeAlpha := uint8(i * 10)
		radius := b.Radius - float32(float32(MaxTrailLength-i)*0.5)
		if fadeAlpha > 0 {
			rl.DrawCircleV(trail.Position, radius, rl.NewColor(245, 245, 220, fadeAlpha))
		}
	}
}

func (b Ball) Rect() rl.Rectangle {
	return rl.NewRectangle(b.X-b.Radius, b.Y-b.Radius, b.Radius*2, b.Radius*2)
}
