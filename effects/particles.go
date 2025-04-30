package effects

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	X, Y   float32
	VX, VY float32
	Life   float32
	Color  rl.Color
	Size   float32
	Active bool
}

func NewParticle(x, y float32, color rl.Color) *Particle {
	return &Particle{
		X:      x,
		Y:      y,
		VX:     rand.Float32()*4 - 1, // -2 to 2
		VY:     rand.Float32() * -2,  // upward burst
		Life:   rand.Float32()*0.5 + 0.1,
		Active: true,
		Color:  color,
	}
}

func (p *Particle) Update(dt float32) {
	if !p.Active {
		return
	}
	p.X += p.VX
	p.Y += p.VY
	p.VY += 0.2 // gravity
	p.Life -= dt
	if p.Life <= 0 {
		p.Active = false
	}
}

func (p *Particle) Draw() {
	if p.Active {
		fadeColor := rl.Fade(p.Color, p.Life) // fades as life decreases
		rl.DrawCircleV(rl.NewVector2(p.X, p.Y), 3, fadeColor)
	}
}
