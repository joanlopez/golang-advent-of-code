package panel

import "fmt"

const (
	// colors
	black = 0
	white = 1

	defaultColor = black
)

type Color int

func (pc Color) IsWhite() bool {
	return pc == white
}

func (pc Color) IsBlack() bool {
	return pc == black
}

type ID string

type Panel struct {
	X int
	Y int
}

func New(x, y int) Panel {
	return Panel{X: x, Y: y}
}

func (p Panel) ID() ID {
	return ID(fmt.Sprintf("(%d,%d)", p.X, p.Y))
}

type Map struct {
	panels map[ID]Color
}

func NewMap() *Map {
	return &Map{panels: make(map[ID]Color)}
}

func (m *Map) ColorAt(id ID) Color {
	if color, ok := m.panels[id]; ok {
		return color
	}

	return defaultColor
}

func (m *Map) SetColorAt(id ID, color Color) {
	m.panels[id] = color
}

func (m *Map) Length() int {
	return len(m.panels)
}
