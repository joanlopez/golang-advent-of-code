package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input-1"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var highestID int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := NewSeat(scanner.Text())
		if s.ID() > highestID {
			highestID = s.ID()
		}
	}

	fmt.Println(highestID)
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
