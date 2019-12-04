package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	WIDTH  = 20
	HEIGHT = 20

	STARTX = 1
	STARTY = HEIGHT - 2

	SEP = ","

	DOT    = 0
	CIRCLE = 1
	VLINE  = 2
	HLINE  = 3
	SUM    = 4
	CROSS  = 5

	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"

	inputFile = "input.txt"
)

type Pos struct {
	x int
	y int
}

func main() {
	//wireA := make([]uint8, WIDTH*HEIGHT)
	//wireB := make([]uint8, WIDTH*HEIGHT)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	input1 := scanner.Text()
	//input1 := "R8,U5,L5,D3"
	input1Moves := strings.Split(input1, SEP)

	scanner.Scan()
	input2 := scanner.Text()
	//input2 := "U7,R6,D4,L4"
	input2Moves := strings.Split(input2, SEP)

	//drawWire(input1Moves, wireA, STARTX, STARTY)
	//drawWire(input2Moves, wireB, STARTX, STARTY)

	pathA := getPath(input1Moves, STARTX, STARTY)
	pathB := getPath(input2Moves, STARTX, STARTY)

	//printSlice(wireA)
	//printSlice(wireB)

	//fmt.Println(findNearestMatchDistance(wireA, wireB, STARTX, STARTY))

	//fmt.Println(findNearestPathDistance(pathA, pathB, Pos{x: STARTX, y: STARTY}))

	fmt.Println(findShortestPathDistance(pathA, pathB, Pos{x: STARTX, y: STARTY}))
}

func slicePos(row, col int) int {
	return row*WIDTH + col
}

func calculateDistance(x1, x2, y1, y2 int) int {
	return int(math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
}

func calculatePosDistance(p1, p2 Pos) int {
	return calculateDistance(p1.x, p2.x, p1.y, p2.y)
}

func areEqual(p1, p2 Pos) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func printSlice(cells []uint8) {
	for row := 0; row < HEIGHT; row++ {
		line := make([]string, WIDTH)

		for col := 0; col < WIDTH; col++ {
			line[col] = drawCell(cells[slicePos(row, col)])
		}
		fmt.Println(line)
	}
}

func drawCell(n uint8) (s string) {
	switch n {
	case DOT:
		s = "."
	case CIRCLE:
		s = "O"
	case VLINE:
		s = "|"
	case HLINE:
		s = "-"
	case SUM:
		s = "+"
	case CROSS:
		s = "X"
	}
	return
}

func drawWire(input1Moves []string, wireA []uint8, startX, startY int) {
	currX, currY := startX, startY
	wireA[slicePos(currY, currX)] = CIRCLE

	for _, move := range input1Moves {

		moveDir := string(move[0])
		moveLen, _ := strconv.Atoi(move[1:])

		if moveDir == RIGHT {
			for x := currX + 1; x < currX+moveLen; x++ {
				wireA[slicePos(currY, x)] = HLINE
			}
			currX += moveLen
			wireA[slicePos(currY, currX)] = SUM
		}

		if moveDir == LEFT {
			for x := currX - 1; x > currX-moveLen; x-- {
				wireA[slicePos(currY, x)] = HLINE
			}
			currX -= moveLen
			wireA[slicePos(currY, currX)] = SUM
		}

		if moveDir == DOWN {
			for y := currY + 1; y < currY+moveLen; y++ {
				wireA[slicePos(y, currX)] = VLINE
			}
			currY += moveLen
			wireA[slicePos(currY, currX)] = SUM
		}

		if moveDir == UP {
			for y := currY - 1; y > currY-moveLen; y-- {
				wireA[slicePos(y, currX)] = VLINE
			}
			currY -= moveLen
			wireA[slicePos(currY, currX)] = SUM
		}
	}
}

func getPath(input1Moves []string, startX, startY int) (path []Pos) {
	currX, currY := startX, startY

	for _, move := range input1Moves {

		moveDir := string(move[0])
		moveLen, _ := strconv.Atoi(move[1:])

		if moveDir == RIGHT {
			for x := currX + 1; x <= currX+moveLen; x++ {
				path = append(path, Pos{x: x, y: currY})
			}
			currX += moveLen
		}

		if moveDir == LEFT {
			for x := currX - 1; x >= currX-moveLen; x-- {
				path = append(path, Pos{x: x, y: currY})
			}
			currX -= moveLen
		}

		if moveDir == DOWN {
			for y := currY + 1; y <= currY+moveLen; y++ {
				path = append(path, Pos{x: currX, y: y})
			}
			currY += moveLen
		}

		if moveDir == UP {
			for y := currY - 1; y >= currY-moveLen; y-- {
				path = append(path, Pos{x: currX, y: y})
			}
			currY -= moveLen
		}
	}

	return
}

func findNearestMatchDistance(wireA, wireB []uint8, toX, toY int) int {
	minDistance := WIDTH * HEIGHT

	for row := 0; row < HEIGHT; row++ {
		for col := 0; col < WIDTH; col++ {
			a := wireA[slicePos(row, col)]
			b := wireB[slicePos(row, col)]

			if (a == VLINE || a == HLINE || a == SUM) &&
				(b == VLINE || b == HLINE || b == SUM) {

				distance := calculateDistance(col, toX, row, toY)
				if minDistance > distance {
					minDistance = distance
				}
			}
		}
	}

	return minDistance
}

func findNearestPathDistance(pathA, pathB []Pos, startPos Pos) int {
	minDistance := math.MaxInt64

	for _, posA := range pathA {
		for _, posB := range pathB {
			if areEqual(posA, posB) {
				distance := calculatePosDistance(posA, startPos)
				if minDistance > distance {
					minDistance = distance
				}
			}
		}
	}

	return minDistance
}

func findShortestPathDistance(pathA, pathB []Pos, startPos Pos) int {
	minSteps := math.MaxInt64

	for x, posA := range pathA {
		for y, posB := range pathB {
			if areEqual(posA, posB) {
				steps := (x+1) + (y+1)
				if minSteps > steps {
					minSteps = steps
				}
			}
		}
	}

	return minSteps
}
