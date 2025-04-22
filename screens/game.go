package screens

import rl "github.com/gen2brain/raylib-go/raylib"

type State int

const (
	Unknown State = iota
	Playing
	Paused
	Finished
	Menu
)

type Game struct {
	State State
	Font  rl.Font
}
