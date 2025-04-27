package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	X, Y   float32
	Width  float32
	Height float32
	Speed  float32
	Color  rl.Color
}

func NewPlayer(x, y float32) Player {
	return Player{
		X:      x,
		Y:      y,
		Width:  100,
		Height: 20,
		Speed:  5,
		Color:  rl.RayWhite,
	}
}

func (p Player) GetRect() rl.Rectangle {
	return rl.NewRectangle(p.X, p.Y, p.Width, p.Height)
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyLeft) {
		if p.X == 0 {
			p.X = 0
		} else if rl.IsKeyDown(rl.KeySpace) {
			p.X -= p.Speed * 2
		} else {
			p.X -= p.Speed
		}
	}
	if rl.IsKeyDown(rl.KeyRight) {
		if p.X+p.Width >= float32(rl.GetScreenWidth()) {
			p.X = float32(rl.GetScreenWidth()) - p.Width
		} else if rl.IsKeyDown(rl.KeySpace) {
			p.X += p.Speed * 2
		} else {
			p.X += p.Speed
		}
	}
}

func (p Player) Draw() {
	rl.DrawRectangleRounded(
		rl.Rectangle{
			X:      p.X,
			Y:      p.Y,
			Width:  p.Width,
			Height: p.Height,
		},
		0.7,
		2,
		p.Color,
	)
}
