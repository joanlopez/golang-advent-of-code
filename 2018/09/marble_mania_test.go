package marble

import (
	"fmt"
	"testing"
)

type marbleManiaTest struct {
	numPlayers       int
	maxMarbles       int
	expectedMaxScore int
}

var tests = []marbleManiaTest{
	{9, 25, 32},
	{10, 1618, 8317},
	{13, 7999, 146373},
	{17, 1104, 2764},
	{21, 6111, 54718},
	{30, 5807, 37305},
}

func TestMarbleMania(t *testing.T) {
	for _, tc := range tests {
		testName := fmt.Sprintf("marble mania with %v players and %v as last marble", tc.numPlayers, tc.maxMarbles)
		t.Run(testName, func(t *testing.T) {
			g := NewGame(tc.numPlayers, tc.maxMarbles); g.Simulate()
			if tc.expectedMaxScore != g.MaxScore {
				t.Errorf("incorrect max score: %v, expected: %v", g.MaxScore, tc.expectedMaxScore)
			}
		})
	}
}
