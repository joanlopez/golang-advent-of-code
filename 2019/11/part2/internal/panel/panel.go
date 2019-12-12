package panel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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
	return ID(fmt.Sprintf("%d,%d", p.X, p.Y))
}

type Map struct {
	panels map[ID]Color
}

func NewMap() *Map {
	return &Map{panels: make(map[ID]Color)}
}

func (m *Map) ColorAt(id ID) Color {
	// Special case for the given statement
	if id == New(0, 0).ID() {
		return white
	}

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

func (m *Map) Print() {
	topX, topY := math.MinInt64, math.MinInt64
	bottomX, bottomY := math.MaxInt64, math.MaxInt64

	for k := range m.panels {
		xy := strings.Split(string(k), ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])

		if x > topX {
			topX = x
		}

		if y > topY {
			topY = y
		}

		if x < bottomX {
			bottomX = x
		}

		if y < bottomY {
			bottomY = y
		}
	}

	// calculate bounds
	height := abs(bottomY - topY)
	width := abs(bottomX - topX)

	// create matrix
	cells := make([][]Color, height+1)
	for i := range cells {
		cells[i] = make([]Color, width+1)
	}

	// fill matrix
	for k, v := range m.panels {
		xy := strings.Split(string(k), ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])

		absBottomX := abs(bottomX)
		absBottomY := abs(bottomY)
		cells[y+absBottomY][x+absBottomX] = v
	}

	// print matrix
	for i := range cells {
		for j := range cells[height-i] {
			if cells[height-i][j].IsBlack() {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

// auxiliary function
func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
