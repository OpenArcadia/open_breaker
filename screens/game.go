package screens

import (
	"open_breaker/effects"
	"open_breaker/entity"
	"open_breaker/utility"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScreen struct {
	StartTime    time.Time
	CurrentLevel LevelName
	Font         *rl.Font
	Player       *entity.Player
	Bricks       []*entity.Brick
	Ball         *entity.Ball
	BounceSound  *rl.Sound
	BreakSound   *rl.Sound
	Particles    []*effects.Particle
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

	breakSound := rl.LoadSound(utility.LoadAssetFrom("music/break.ogg"))
	rl.SetSoundVolume(breakSound, 1.0)
	bounceSound := rl.LoadSound(utility.LoadAssetFrom("music/bounce.ogg"))
	rl.SetSoundVolume(bounceSound, 1.0)

	g.BounceSound = &bounceSound
	g.BreakSound = &breakSound
	g.StartTime = time.Now()

	g.Ball = entity.NewBall(400, 300, &bounceSound)
	g.Player = entity.NewPlayer(350, float32(rl.GetScreenHeight()-100))

	bricks := []*entity.Brick{}
	switch g.CurrentLevel {
	case LEVEL_ONE:
		// 10 columns, 3 rows
		for i := 0; i < 10; i++ {
			for j := 0; j < 3; j++ {
				bricks = append(bricks, entity.NewBrick(70, 30, 150+float32(i)*80, 80+float32(j)*40, &breakSound, true))
			}
		}

	case LEVEL_TWO:
		for i := 0; i < 12; i++ {
			for j := 0; j < 5; j++ {
				breakable := !(j == 0 && i%3 == 0) // every 3rd block in top row is unbreakable
				bricks = append(bricks, entity.NewBrick(70, 30, 80+float32(i)*80, 50+float32(j)*40, &breakSound, breakable))
			}
		}

	case LEVEL_THREE:
		for i := 0; i < 13; i++ {
			for j := 0; j < 6; j++ {
				breakable := !(j == 5 && i != 6)
				bricks = append(bricks, entity.NewBrick(70, 30, 30+float32(i)*80, 30+float32(j)*40, &breakSound, breakable))
			}
		}

	case LEVEL_FOUR:
		for i := 0; i < 13; i++ {
			for j := 0; j < 6; j++ {
				breakable := (!(j == 5 && i != 6) && !((i == 5 || i == 7) && j >= 2))
				bricks = append(bricks, entity.NewBrick(70, 30, 30+float32(i)*80, 30+float32(j)*40, &breakSound, breakable))
			}
		}
	case LEVEL_FIVE:
		for i := 0; i < 21; i++ {
			for j := 0; j < 7; j++ {
				if j == 4 {
					continue
				}
				x := 30 + float32(i)*50
				y := 30 + float32(j)*30

				unbreakableColumn := j == 6 || j == 3
				forceBreakRow := i == 11

				breakable := !(unbreakableColumn && !forceBreakRow)

				bricks = append(bricks, entity.NewBrick(40, 20, x, y, &breakSound, breakable))
			}
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

	if len(Filter(g.Bricks, func(b *entity.Brick) bool {
		return b.Breakable
	})) == 0 {
		nextScreenMap := map[LevelName]LevelName{
			LEVEL_ONE:   LEVEL_TWO,
			LEVEL_TWO:   LEVEL_THREE,
			LEVEL_THREE: LEVEL_FOUR,
			LEVEL_FOUR:  LEVEL_FIVE,
		}
		ChangeScreen(&FinishScreen{
			Font:      g.Font,
			StartTime: g.StartTime,
			EndTime:   time.Now(),
			NextLevel: nextScreenMap[g.CurrentLevel],
		})
	}

	if g.Ball.Y >= float32(rl.GetScreenHeight()) {
		ChangeScreen(&GameOverScreen{
			Font:      g.Font,
			FromLevel: g.CurrentLevel,
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
