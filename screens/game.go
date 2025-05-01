package screens

import (
	"open_breaker/effects"
	"open_breaker/entity"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State int

type GameScreen struct {
	State        State
	StartTime    float64
	EndTime      float64
	CurrentLevel LevelName
	Font         *rl.Font
	Player       *entity.Player
	Bricks       []*entity.Brick
	Ball         *entity.Ball
	BounceSound  *rl.Sound
	BreakSound   *rl.Sound

	Particles []*effects.Particle
}

func Filter[T any](input []T, test func(T) bool) []T {
	var result []T
	for _, v := range input {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

func (g *GameScreen) Create() {
	_, isFlatpak := os.LookupEnv("container")

	var basePath string
	if isFlatpak {
		basePath = "/app/bin/assets/"
	} else {
		basePath = "assets/"
	}
	breakSound := rl.LoadSound(basePath + "music/break.ogg")
	rl.SetSoundVolume(breakSound, 1.0)
	bounceSound := rl.LoadSound(basePath + "music/bounce.ogg")
	rl.SetSoundVolume(bounceSound, 1.0)

	g.BounceSound = &bounceSound
	g.BreakSound = &breakSound

	g.Ball = entity.NewBall(400, 300, &bounceSound)
	g.Player = entity.NewPlayer(350, float32(rl.GetScreenHeight()-100))

	bricks := []*entity.Brick{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			bricks = append(bricks, entity.NewBrick(150+float32(i)*80, 50+float32(j)*40, &breakSound))
		}
	}
	g.Bricks = bricks
}

func (g *GameScreen) Render() {

	g.Player.Update()
	g.Ball.Update(*g.Player)

	g.Bricks = Filter(g.Bricks, func(b *entity.Brick) bool {
		return b.Visible
	})

	if len(g.Bricks) == 0 {
		ChangeScreen(&FinishScreen{
			Font: g.Font,
		})
	}

	if g.Ball.Y >= float32(rl.GetScreenHeight()) {
		ChangeScreen(&GameOverScreen{
			Font: g.Font,
		})
	}
	dt := rl.GetFrameTime()
	for i := range g.Particles {
		g.Particles[i].Update(dt)
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(10, 15, 25, 255))

	g.Player.Draw()
	g.Ball.Draw()
	for i := range g.Bricks {
		b := g.Bricks[i]
		b.Draw()
		if !b.Update(g.Ball) {
			for j := 0; j < 10; j++ {
				pX := b.X + b.Width/2
				pY := b.Y + b.Height/2
				g.Particles = append(g.Particles, effects.NewParticle(pX, pY, b.PrimaryColor))
			}
		}
	}

	for _, p := range g.Particles {
		p.Draw()
	}

	rl.EndDrawing()

	g.Particles = Filter(g.Particles, func(p *effects.Particle) bool {
		return p.Active
	})
}

func (g *GameScreen) Dispose() {
	rl.UnloadSound(*g.BounceSound)
	rl.UnloadSound(*g.BreakSound)
	g.BounceSound = nil
	g.BreakSound = nil
}
