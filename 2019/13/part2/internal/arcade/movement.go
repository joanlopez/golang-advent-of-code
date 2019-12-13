package arcade

const (
	xpos = 0
	ypos = 1
	kind = 2
)

type Mode int

func (m *Mode) IsXPos() bool {
	return *m == xpos
}

func (m *Mode) IsYPos() bool {
	return *m == ypos
}

func (m *Mode) IsKind() bool {
	return *m == kind
}

func (m *Mode) Next() {
	*m = (*m + 1) % 3
}

const (
	// joystick moves
	neutral = 0
	left    = -1
	right   = 1
)

type JoystickMove int

func (m JoystickMove) IsNeutral() bool {
	return m == neutral
}

func (m JoystickMove) IsLeft() bool {
	return m == left
}

func (m JoystickMove) IsRight() bool {
	return m == right
}
