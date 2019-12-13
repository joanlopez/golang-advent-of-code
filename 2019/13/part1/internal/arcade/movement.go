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
