package robot

import (
	"../intcode"
	"../panel"
)

type Robot struct {
	panelMap *panel.Map

	program *intcode.Program

	mode      *Mode
	direction Direction
	currPanel panel.Panel
}

func New(program *intcode.Program) *Robot {
	defaultMode := Mode(draw)
	return &Robot{
		panelMap: panel.NewMap(),
		program:  program,
		mode:     &defaultMode,
	}
}

func (r *Robot) Run() {
	io := NewIO()
	go r.program.Execute(io)

exec:
	for {
		msg := <-io.C
		switch {
		case msg.IsHalt():
			break exec
		case msg.IsInput():
			msg.C <- int(r.panelMap.ColorAt(r.currPanel.ID()))
		case msg.IsOutput():
			output := <-msg.C

			if r.mode.IsDraw() {
				r.panelMap.SetColorAt(r.currPanel.ID(), panel.Color(output))
			}

			if r.mode.IsMove() {
				r.direction = movements[output][r.direction]
				r.currPanel = r.nextPanel()
			}

			r.mode.Switch()
		}
	}
}

func (r *Robot) nextPanel() panel.Panel {
	if r.direction.IsUp() {
		return panel.New(r.currPanel.X, r.currPanel.Y+1)
	}

	if r.direction.IsDown() {
		return panel.New(r.currPanel.X, r.currPanel.Y-1)
	}

	if r.direction.IsLeft() {
		return panel.New(r.currPanel.X-1, r.currPanel.Y)
	}

	if r.direction.IsRight() {
		return panel.New(r.currPanel.X+1, r.currPanel.Y)
	}

	// default
	return r.currPanel
}

func (r *Robot) VisitedPanels() int {
	return r.panelMap.Length()
}
