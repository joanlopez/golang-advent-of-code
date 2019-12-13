package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./internal/arcade"
	"./internal/intcode"
	"./internal/tile"
)

const (
	inputFile = "input.txt"
)

func main() {
	// read input
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// initializing
	programDefinition := strings.Split(string(content), ",")
	program, err := intcode.NewProgram(programDefinition)
	if err != nil {
		log.Fatal(err)
	}

	// and running the arcade cabinet with the given program
	cabinet := arcade.NewCabinet(program)
	cabinet.Run()

	// returning the total number of block tiles
	fmt.Printf("total block tiles: %d\n", cabinet.CountTilesByfType(tile.Kind(tile.Block)))
}
