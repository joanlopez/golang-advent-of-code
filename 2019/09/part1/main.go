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
	BASE = 9
	HALT = 99

	// modes
	UNKNOWN   = 0
	POSITION  = 0
	IMMEDIATE = 1
	RELATIVE  = 2

	// default
	INPUT = 1

	inputFile = "input.txt"
)

type Instruction struct {
	currPos  int
	currBase int
	opCode   int
	modes    []int
	args     []int
	program  map[int]int
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

func (i Instruction) IsBase() bool {
	return i.opCode == BASE
}

func (i Instruction) IsHalt() bool {
	return i.opCode == HALT
}

func (i Instruction) Param(param int) int {
	if i.modes[param-1] == IMMEDIATE {
		return i.args[param-1]
	}

	if i.modes[param-1] == RELATIVE {
		return i.program[i.currBase+i.args[param-1]]
	}

	return i.program[i.args[param-1]]
}

func (i Instruction) MemPos(param int) int {
	if i.modes[param-1] == POSITION {
		return i.args[param-1]
	} else { // relative
		return i.currBase + i.args[param-1]
	}
}

func (i Instruction) Length() int {
	return len(i.args) + 1
}

// Execute returns the next position
func (i Instruction) Execute() (int, int) {
	if i.IsSum() || i.IsMul() || i.IsLessThan() || i.IsEqual() {
		x := i.Param(1)
		y := i.Param(2)
		res := i.MemPos(3)

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

		return i.currPos + 4, i.currBase
	}

	if i.IsJumpFalse() || i.IsJumpTrue() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsJumpTrue()) {
			return y, i.currBase
		} else {
			return i.currPos + 3, i.currBase
		}
	}

	if i.IsLessThan() || i.IsEqual() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsEqual()) {
			return y, i.currBase
		} else {
			return i.currPos + 3, i.currBase
		}
	}

	if i.IsBase() {
		return i.currPos + 2, i.currBase + i.Param(1)
	}

	if i.IsIn() {
		i.program[i.MemPos(1)] = INPUT
	} else {
		fmt.Println(i.Param(1))
	}
	return i.currPos + 2, i.currBase
}

func NewInstruction(currPos, currBase int, program map[int]int) Instruction {
	op := program[currPos]

	i := Instruction{
		currPos:  currPos,
		currBase: currBase,
		opCode:   op % 100,
		modes:    []int{(op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10},
		program:  program,
	}

	if i.IsSum() || i.IsMul() || i.IsLessThan() || i.IsEqual() {
		i.args = []int{program[currPos+1], program[currPos+2], program[currPos+3]}
	} else if i.IsJumpTrue() || i.IsJumpFalse() {
		i.args = []int{program[currPos+1], program[currPos+2]}
	} else if i.IsIn() || i.IsOut() || i.IsBase() {
		i.args = []int{program[currPos+1]}
	}

	return i
}

func doExecution(program map[int]int) map[int]int {
	currPos, currBase := 0, 0

	for {
		i := NewInstruction(currPos, currBase, program)
		if i.IsHalt() {
			break
		}

		currPos, currBase = i.Execute()
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

		program := make(map[int]int, len(elems))
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
