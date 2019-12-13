package arcade

import (
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
			panic("input op not expected")
		case msg.IsOutput():
			output := <-msg.C

			if r.mode.IsXPos() {
				r.lastX = output
			}

			if r.mode.IsYPos() {
				r.lastY = output
			}

			if r.mode.IsKind() {
				newPanel := tile.New(r.lastX, r.lastY)
				r.tileMap.SetKindAt(newPanel.ID(), tile.Kind(output))
			}

			r.mode.Next()
		}
	}
}

func (r *Cabinet) CountTilesByfType(kind tile.Kind) int {
	return r.tileMap.CountByKind(kind)
}
