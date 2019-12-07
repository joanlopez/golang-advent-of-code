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
	// opCode codes
	SUM  = 1
	MUL  = 2
	IN   = 3
	OUT  = 4
	JT   = 5
	JF   = 6
	LT   = 7
	EQ   = 8
	HALT = 99

	// modes
	UNKNOWN   = 0
	POSITION  = 0
	IMMEDIATE = 1

	// default
	INPUT = 5

	search    = 19690720
	inputFile = "input.txt"
)

type Instruction struct {
	currPos int
	opCode  int
	modes   []int
	args    []int
	program []int
}

func (i Instruction) IsSum() bool {
	return i.opCode == SUM
}

func (i Instruction) IsMul() bool {
	return i.opCode == MUL
}

func (i Instruction) IsIn() bool {
	return i.opCode == IN
}

func (i Instruction) IsOut() bool {
	return i.opCode == OUT
}

func (i Instruction) IsJumpTrue() bool {
	return i.opCode == JT
}

func (i Instruction) IsJumpFalse() bool {
	return i.opCode == JF
}

func (i Instruction) IsLessThan() bool {
	return i.opCode == LT
}

func (i Instruction) IsEqual() bool {
	return i.opCode == EQ
}

func (i Instruction) IsHalt() bool {
	return i.opCode == HALT
}

func (i Instruction) Param(param int) int {
	if i.modes[param-1] == IMMEDIATE {
		return i.args[param-1]
	}

	return i.program[i.args[param-1]]
}

func (i Instruction) Length() int {
	return len(i.args) + 1
}

// Execute returns the next position
func (i Instruction) Execute() int {
	if i.IsSum() || i.IsMul() || i.IsLessThan() || i.IsEqual() {
		x := i.Param(1)
		y := i.Param(2)
		res := i.args[2]

		if i.IsSum() {
			i.program[res] = x + y
		} else if i.IsMul() {
			i.program[res] = x * y
		} else {
			if (i.IsLessThan() && x < y) || (i.IsEqual() && x == y) {
				i.program[res] = 1
			} else {
				i.program[res] = 0
			}
		}

		return i.currPos + 4
	}

	if i.IsJumpFalse() || i.IsJumpTrue() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsJumpTrue()) {
			return y
		} else {
			return i.currPos + 3
		}
	}

	if i.IsLessThan() || i.IsEqual() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsEqual()) {
			return y
		} else {
			return i.currPos + 3
		}
	}

	if i.IsIn() {
		x := i.args[0]
		i.program[x] = INPUT
	} else {
		fmt.Println(i.Param(1))
	}
	return i.currPos + 2
}

func NewInstruction(currPos int, program []int) Instruction {
	op := program[currPos]

	i := Instruction{
		currPos: currPos,
		opCode:  op % 100,
		modes:   []int{(op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10},
		program: program,
	}

	if i.IsSum() || i.IsMul() || i.IsLessThan() || i.IsEqual() {
		i.args = []int{program[currPos+1], program[currPos+2], program[currPos+3]}
	} else if i.IsJumpTrue() || i.IsJumpFalse() {
		i.args = []int{program[currPos+1], program[currPos+2]}
	} else if i.IsIn() || i.IsOut() {
		i.args = []int{program[currPos+1]}
	}

	return i
}

func doExecution(program []int) []int {
	currPos := 0

	for {
		i := NewInstruction(currPos, program)
		if i.IsHalt() {
			break
		}

		currPos = i.Execute()
	}

	return program
}

// Please, do a refactor. The entire code is very ugly & makes no sense...
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

		doExecution(program)
	}
}
