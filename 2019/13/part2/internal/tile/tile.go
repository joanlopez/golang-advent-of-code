package tile

import (
	"fmt"
)

const (
	// kinds
	Empty = iota
	Wall
	Block
	Paddle
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

func (t Kind) IsPaddle() bool {
	return t == Paddle
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
	ball   Tile
	paddle Tile
	tiles  map[ID]Kind
}

func NewMap() *Map {
	return &Map{tiles: make(map[ID]Kind)}
}

func (m *Map) Ball() *Tile {
	return &m.ball
}

func (m *Map) Paddle() *Tile {
	return &m.paddle
}

func (m *Map) SetTile(tile Tile, kind Kind) {
	if kind.IsBall() {
		m.ball = tile
	}

	if kind.IsPaddle() {
		m.paddle = tile
	}

	m.tiles[tile.ID()] = kind
}

func (m *Map) KindAt(id ID) Kind {
	if kind, ok := m.tiles[id]; ok {
		return kind
	}

	return Empty
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
