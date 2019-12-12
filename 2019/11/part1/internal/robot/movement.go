package robot

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3

	turnLeft  = 0
	turnRight = 1
)

type Direction int

func (d Direction) IsUp() bool {
	return d == up
}

func (d Direction) IsDown() bool {
	return d == down
}

func (d Direction) IsLeft() bool {
	return d == left
}

func (d Direction) IsRight() bool {
	return d == right
}

const (
	draw = 0
	move = 1
)

type Mode int

func (m *Mode) IsDraw() bool {
	return *m == draw
}

func (m *Mode) IsMove() bool {
	return *m == move
}

func (m *Mode) Switch() {
	if *m == draw {
		*m = move
	} else {
		*m = draw
	}
}

var movements = map[int]map[Direction]Direction{
	turnLeft:  {up: left, left: down, down: right, right: up},
	turnRight: {up: right, right: down, down: left, left: up},
}
