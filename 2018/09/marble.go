package marble

type marble struct {
	prev *marble
	next *marble
	val  int
}

func (m *marble) isSpecial() bool {
	return m.val%23 == 0
}

func initZeroMarble() (m *marble) {
	m = &marble{val: 0}
	m.prev = m
	m.next = m
	return
}
