package screens

import rl "github.com/gen2brain/raylib-go/raylib"

type State int

const (
	Unknown State = iota
	Playing
	Paused
	Finished
	Menu
	Over
)

type Game struct {
	State     State
	StartTime float64
	EndTime   float64
	Font      rl.Font
}
