package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	inputFile = "input-2"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lowest, highest := math.MaxInt64, math.MinInt64
	seats := make(map[int]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := NewSeat(scanner.Text())
		seats[s.ID()] = struct{}{}

		if lowest > s.ID() {
			lowest = s.ID()
		}

		if highest < s.ID() {
			highest = s.ID()
		}
	}

	for id := lowest + 1; id < highest; id++ {
		if _, ok := seats[id]; !ok {
			fmt.Println(id)
		}
	}
}

const (
	minRow, maxRow       = 0, 127
	minColumn, maxColumn = 0, 7
)

type Seat struct {
	row    int
	column int
}

func NewSeat(raw string) Seat {
	minRow, maxRow := minRow, maxRow
	minColumn, maxColumn := minColumn, maxColumn

	for _, char := range raw {
		switch char {
		case 'F':
			maxRow -= ((maxRow - minRow) / 2) + 1
		case 'B':
			minRow += ((maxRow - minRow) / 2) + 1
		case 'L':
			maxColumn -= ((maxColumn - minColumn) / 2) + 1
		case 'R':
			minColumn += ((maxColumn - minColumn) / 2) + 1
		}
	}

	return Seat{row: minRow, column: minColumn}
}

func (s Seat) ID() int {
	numColumns := maxColumn + 1
	return s.row*numColumns + s.column
}
