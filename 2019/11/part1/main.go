package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./internal/intcode"
	"./internal/robot"
)

const (
	inputFile = "input.txt"

	separator = ","
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

	// and running a robot with the given program
	programmedRobot := robot.New(program)
	programmedRobot.Run()

	// returning the total number of visited panels
	fmt.Printf("the programmedRobot visited %d panels\n", programmedRobot.VisitedPanels())
}
