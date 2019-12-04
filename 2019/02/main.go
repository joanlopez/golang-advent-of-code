package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// op codes
	ADD = 1
	MUL = 2
	HALT = 99

	search = 19690720
	inputFile = "input.txt"
)

func doExecution(program []int) []int {
	currPos := 0
	op := program[currPos]

	for op != HALT {
		op1, op2, op3 := program[currPos+1], program[currPos+2], program[currPos+3]

		var res int
		if op == ADD {
			res = program[op1] + program[op2]
		} else if op == MUL {
			res = program[op1] * program[op2]
		}
		program[op3] = res

		currPos += 4
		op = program[currPos]
	}

	return program
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), ",")

		program := make([]int, len(elems))
		for i := range elems {
			n, err := strconv.Atoi(elems[i])
			if err != nil {
				log.Fatal(err)
			}

			program[i] = n
		}

		// These values were found by investigating how the final output changes when
		// the noun and the verb are incremented by one unit.
		for noun := 50; noun < 60; noun++ {
			for verb := 80; verb < 90; verb++ {

				programCopy := make([]int, len(program))
				copy(programCopy, program)

				programCopy[1] = noun
				programCopy[2] = verb

				output := doExecution(programCopy)[0]

				if output == search {
					fmt.Printf("The final execution with (%d, %d) was: %v\n", verb, noun, output)
				}
			}



		}
	}
}
