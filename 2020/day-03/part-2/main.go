package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input-2"
)

func main() {
	cellMap, err := readMapFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	steps := []Step{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	fmt.Println(cellMap.totalTrees(steps))
}

type Step struct {
	width  int
	height int
}

type CellMap struct {
	cells []string
}

func (m CellMap) totalTrees(steps []Step) (result int64) {
	if len(steps) == 0 {
		return
	}

	result += m.traverseAndCountTrees(steps[0])

	for i := 1; i < len(steps); i++ {
		result *= m.traverseAndCountTrees(steps[i])
	}

	return
}

func (m CellMap) traverseAndCountTrees(step Step) (count int64) {
	x, y := 0, 0

	// until reach bottom-most row
	for y < len(m.cells)-1 {
		if x+step.width > len(m.cells[y])-1 {
			x = (x + step.width) % len(m.cells[y])
		} else {
			x += step.width
		}

		y += step.height

		if m.cells[y][x] == '#' {
			count++
		}
	}

	return count

}

func readMapFromFile(filename string) (CellMap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return CellMap{}, err
	}

	defer func() {
		err = file.Close()
	}()

	var cells []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cells = append(cells, line)
	}

	return CellMap{cells: cells}, scanner.Err()
}
