package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input-1"

	stepWidth = 3
)

func main() {
	cellMap, err := readMapFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cellMap.traverseAndCountTrees())
}

type CellMap struct {
	cells []string
}

func (m CellMap) traverseAndCountTrees() (count int64) {
	x, y := 0, 0

	// until reach bottom-most row
	for y < len(m.cells)-1 {
		if x+stepWidth > len(m.cells[y])-1 {
			x = (x + stepWidth) % len(m.cells[y])
		} else {
			x += stepWidth
		}

		y++

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
