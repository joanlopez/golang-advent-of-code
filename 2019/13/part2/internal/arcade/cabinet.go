package arcade

import (
	"fmt"

	"../intcode"
	"../tile"
)

type Cabinet struct {
	tileMap *tile.Map

	program *intcode.Program

	mode *Mode

	lastX int
	lastY int
}

func NewCabinet(program *intcode.Program) *Cabinet {
	defaultMode := Mode(xpos)
	return &Cabinet{
		tileMap: tile.NewMap(),
		program: program,
		mode:    &defaultMode,
	}
}

func (r *Cabinet) Run() {
	io := NewIO()
	go r.program.Execute(io)

exec:
	for {
		msg := <-io.C
		switch {
		case msg.IsHalt():
			break exec
		case msg.IsInput():
			msg.C <- int(r.joystickMove())
		case msg.IsOutput():
			output := <-msg.C

			if r.mode.IsXPos() {
				r.lastX = output
			}

			if r.mode.IsYPos() {
				r.lastY = output
			}

			if r.mode.IsKind() {
				if r.isOnPrintScore() {
					fmt.Printf("the current score is: %d\n", output)
					fmt.Printf("remaining blocks: %d\n", r.tileMap.CountByKind(tile.Block))
				} else {
					newPanel := tile.New(r.lastX, r.lastY)
					r.tileMap.SetTile(newPanel, tile.Kind(output))
				}
			}

			r.mode.Next()
		}
	}
}

func (r *Cabinet) CountTilesByfType(kind tile.Kind) int {
	return r.tileMap.CountByKind(kind)
}

func (r *Cabinet) isOnPrintScore() bool {
	return r.lastX == -1 && r.lastY == 0
}

func (r *Cabinet) joystickMove() JoystickMove {
	if r.tileMap.Paddle().X < r.tileMap.Ball().X {
		r.tileMap.Paddle().X++
		return JoystickMove(right)
	}
	if r.tileMap.Paddle().X > r.tileMap.Ball().X {
		r.tileMap.Paddle().X--
		return JoystickMove(left)
	}
	return JoystickMove(neutral)
}
