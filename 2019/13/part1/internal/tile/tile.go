package tile

import (
	"fmt"
)

const (
	// kinds
	Empty = iota
	Wall
	Block
	Horizontal
	Ball
)

type Kind int

func (t Kind) IsEmpty() bool {
	return t == Empty
}

func (t Kind) IsWall() bool {
	return t == Wall
}

func (t Kind) IsBlock() bool {
	return t == Block
}

func (t Kind) IsHorizontal() bool {
	return t == Horizontal
}

func (t Kind) IsBall() bool {
	return t == Ball
}

type ID string

type Tile struct {
	X int
	Y int
}

func New(x, y int) Tile {
	return Tile{X: x, Y: y}
}

func (p Tile) ID() ID {
	return ID(fmt.Sprintf("%d,%d", p.X, p.Y))
}

type Map struct {
	tiles map[ID]Kind
}

func NewMap() *Map {
	return &Map{tiles: make(map[ID]Kind)}
}

func (m *Map) KindAt(id ID) Kind {
	if kind, ok := m.tiles[id]; ok {
		return kind
	}

	return Empty
}

func (m *Map) SetKindAt(id ID, kind Kind) {
	m.tiles[id] = kind
}

func (m *Map) Length() int {
	return len(m.tiles)
}

func (m *Map) CountByKind(Kind Kind) (count int) {
	for i := range m.tiles {
		if m.tiles[i] == Kind {
			count++
		}
	}

	return
}
