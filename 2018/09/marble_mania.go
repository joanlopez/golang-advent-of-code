package marble_mania

const numOfSpecialMovements = 7

type Game struct {
	NumPlayers int
	MaxMarbles int
	scores     []int
	MaxScore   int
	currMarble *marble
	nextMarble *marble
	play       int
}

func NewGame(numPlayers, maxMarbles int) *Game {
	marble := initZeroMarble()
	return &Game{
		NumPlayers: numPlayers,
		MaxMarbles: maxMarbles,
		scores:     make([]int, numPlayers),
		currMarble: marble,
		nextMarble: marble,
	}
}

func (g *Game) Simulate() {
	for !g.isFinished() {
		g.nextPlay()

		if g.nextMarble.isSpecial() {
			g.doSpecial()
			continue
		}

		// Put the marble on the proper place
		g.nextMarble.next = g.currMarble.next.next
		g.currMarble.next.next.prev = g.nextMarble
		g.currMarble.next.next = g.nextMarble
		g.nextMarble.prev = g.currMarble.next
	}
}

func (g *Game) isFinished() bool {
	return g.nextMarble.val == g.MaxMarbles
}

func (g *Game) nextPlay() {
	g.play++

	if !g.nextMarble.isSpecial() {
		g.currMarble = g.nextMarble
	}

	g.nextMarble = &marble{val: g.nextMarble.val + 1}
}

func (g *Game) doSpecial() {
	// Getting the marble to remove
	for i := 0; i <= numOfSpecialMovements; i++ {
		g.currMarble = g.currMarble.prev
	}

	// Updating score
	earnedPoints := g.nextMarble.val + g.currMarble.next.val
	g.scores[g.play%g.NumPlayers] += earnedPoints

	// Updating max score
	if g.currentPlayerScore() > g.MaxScore {
		g.MaxScore = g.currentPlayerScore()
	}

	// Removing marble & updating current marble
	g.currMarble.next = g.currMarble.next.next
	g.currMarble.next.prev = g.currMarble
	g.currMarble = g.currMarble.next
}

func (g *Game) currentPlayerScore() int {
	return g.scores[g.play%g.NumPlayers]
}