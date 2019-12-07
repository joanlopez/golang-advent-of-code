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
	INPUT = 0

	search    = 19690720
	inputFile = "input.txt"
)

type Instruction struct {
	currPos int
	opCode  int
	modes   []int
	args    []int
	program []int
	input   int
}

func (i Instruction) IsSum() bool {
	return i.opCode == SUM
}

func (i Instruction) IsMul() bool {
	return i.opCode == MUL
}

func (i Instruction) IsInput() bool {
	return i.opCode == IN
}

func (i Instruction) IsOutput() bool {
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
func (i Instruction) Execute() (int, *int) {
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

		return i.currPos + 4, nil
	}

	if i.IsJumpFalse() || i.IsJumpTrue() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsJumpTrue()) {
			return y, nil
		} else {
			return i.currPos + 3, nil
		}
	}

	if i.IsLessThan() || i.IsEqual() {
		x := i.Param(1)
		y := i.Param(2)

		if (x == 0 && i.IsJumpFalse()) || (x != 0 && i.IsEqual()) {
			return y, nil
		} else {
			return i.currPos + 3, nil
		}
	}

	if i.IsInput() {
		x := i.args[0]
		i.program[x] = i.input
		return i.currPos + 2, nil
	} else {
		output := i.Param(1)
		return i.currPos + 2, &output
	}
}

func NewInstruction(currPos int, program []int, input int) Instruction {
	op := program[currPos]

	i := Instruction{
		currPos: currPos,
		opCode:  op % 100,
		modes:   []int{(op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10},
		program: program,
		input:   input,
	}

	if i.IsSum() || i.IsMul() || i.IsLessThan() || i.IsEqual() {
		i.args = []int{program[currPos+1], program[currPos+2], program[currPos+3]}
	} else if i.IsJumpTrue() || i.IsJumpFalse() {
		i.args = []int{program[currPos+1], program[currPos+2]}
	} else if i.IsInput() || i.IsOutput() {
		i.args = []int{program[currPos+1]}
	}

	return i
}

type Amplifier struct {
	program []int

	currPos int
	setting int
}

func NewAmplifier(setting int, program []int) *Amplifier {
	return &Amplifier{
		currPos: 0,
		program: program,
		setting: setting,
	}
}

func (a *Amplifier) Run(input int) (int, bool) {
	var output *int
	var currInput int
	var inputs []int

	if a.currPos == 0 {
		inputs = []int{a.setting, input}
	} else {
		inputs = []int{input}
	}

	for {
		input = inputs[currInput]

		i := NewInstruction(a.currPos, a.program, input)

		if i.IsHalt() {
			return 0, true
		}

		if i.IsInput() {
			currInput = (currInput + 1) % len(inputs)
		}

		a.currPos, output = i.Execute()

		if output != nil {
			return *output, false
		}
	}
}

var combinations = [][]int{
	{5, 6, 7, 8, 9},
	{5, 6, 7, 9, 8},
	{5, 6, 8, 7, 9},
	{5, 6, 8, 9, 7},
	{5, 6, 9, 7, 8},
	{5, 6, 9, 8, 7},
	{5, 7, 6, 8, 9},
	{5, 7, 6, 9, 8},
	{5, 7, 8, 6, 9},
	{5, 7, 8, 9, 6},
	{5, 7, 9, 6, 8},
	{5, 7, 9, 8, 6},
	{5, 8, 6, 7, 9},
	{5, 8, 6, 9, 7},
	{5, 8, 7, 6, 9},
	{5, 8, 7, 9, 6},
	{5, 8, 9, 6, 7},
	{5, 8, 9, 7, 6},
	{5, 9, 6, 7, 8},
	{5, 9, 6, 8, 7},
	{5, 9, 7, 6, 8},
	{5, 9, 7, 8, 6},
	{5, 9, 8, 6, 7},
	{5, 9, 8, 7, 6},
	{6, 5, 7, 8, 9},
	{6, 5, 7, 9, 8},
	{6, 5, 8, 7, 9},
	{6, 5, 8, 9, 7},
	{6, 5, 9, 7, 8},
	{6, 5, 9, 8, 7},
	{6, 7, 5, 8, 9},
	{6, 7, 5, 9, 8},
	{6, 7, 8, 5, 9},
	{6, 7, 8, 9, 5},
	{6, 7, 9, 5, 8},
	{6, 7, 9, 8, 5},
	{6, 8, 5, 7, 9},
	{6, 8, 5, 9, 7},
	{6, 8, 7, 5, 9},
	{6, 8, 7, 9, 5},
	{6, 8, 9, 5, 7},
	{6, 8, 9, 7, 5},
	{6, 9, 5, 7, 8},
	{6, 9, 5, 8, 7},
	{6, 9, 7, 5, 8},
	{6, 9, 7, 8, 5},
	{6, 9, 8, 5, 7},
	{6, 9, 8, 7, 5},
	{7, 5, 6, 8, 9},
	{7, 5, 6, 9, 8},
	{7, 5, 8, 6, 9},
	{7, 5, 8, 9, 6},
	{7, 5, 9, 6, 8},
	{7, 5, 9, 8, 6},
	{7, 6, 5, 8, 9},
	{7, 6, 5, 9, 8},
	{7, 6, 8, 5, 9},
	{7, 6, 8, 9, 5},
	{7, 6, 9, 5, 8},
	{7, 6, 9, 8, 5},
	{7, 8, 5, 6, 9},
	{7, 8, 5, 9, 6},
	{7, 8, 6, 5, 9},
	{7, 8, 6, 9, 5},
	{7, 8, 9, 5, 6},
	{7, 8, 9, 6, 5},
	{7, 9, 5, 6, 8},
	{7, 9, 5, 8, 6},
	{7, 9, 6, 5, 8},
	{7, 9, 6, 8, 5},
	{7, 9, 8, 5, 6},
	{7, 9, 8, 6, 5},
	{8, 5, 6, 7, 9},
	{8, 5, 6, 9, 7},
	{8, 5, 7, 6, 9},
	{8, 5, 7, 9, 6},
	{8, 5, 9, 6, 7},
	{8, 5, 9, 7, 6},
	{8, 6, 5, 7, 9},
	{8, 6, 5, 9, 7},
	{8, 6, 7, 5, 9},
	{8, 6, 7, 9, 5},
	{8, 6, 9, 5, 7},
	{8, 6, 9, 7, 5},
	{8, 7, 5, 6, 9},
	{8, 7, 5, 9, 6},
	{8, 7, 6, 5, 9},
	{8, 7, 6, 9, 5},
	{8, 7, 9, 5, 6},
	{8, 7, 9, 6, 5},
	{8, 9, 5, 6, 7},
	{8, 9, 5, 7, 6},
	{8, 9, 6, 5, 7},
	{8, 9, 6, 7, 5},
	{8, 9, 7, 5, 6},
	{8, 9, 7, 6, 5},
	{9, 5, 6, 7, 8},
	{9, 5, 6, 8, 7},
	{9, 5, 7, 6, 8},
	{9, 5, 7, 8, 6},
	{9, 5, 8, 6, 7},
	{9, 5, 8, 7, 6},
	{9, 6, 5, 7, 8},
	{9, 6, 5, 8, 7},
	{9, 6, 7, 5, 8},
	{9, 6, 7, 8, 5},
	{9, 6, 8, 5, 7},
	{9, 6, 8, 7, 5},
	{9, 7, 5, 6, 8},
	{9, 7, 5, 8, 6},
	{9, 7, 6, 5, 8},
	{9, 7, 6, 8, 5},
	{9, 7, 8, 5, 6},
	{9, 7, 8, 6, 5},
	{9, 8, 5, 6, 7},
	{9, 8, 5, 7, 6},
	{9, 8, 6, 5, 7},
	{9, 8, 6, 7, 5},
	{9, 8, 7, 5, 6},
	{9, 8, 7, 6, 5},
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

		max := math.MinInt64

		for _, phaseSetting := range combinations {
			amplifiers := make([]*Amplifier, 0, 5)

			for _, setting := range phaseSetting {
				programCopy := make([]int, len(program))
				copy(programCopy, program)

				amplifiers = append(amplifiers, NewAmplifier(setting, programCopy))
			}

			var halted bool
			var amplifier int
			var output, lastOutput int

			for {
				output, halted = amplifiers[amplifier].Run(output)

				if halted {
					break
				}

				lastOutput = output
				amplifier = (amplifier + 1) % len(amplifiers)
			}

			if lastOutput > max {
				max = lastOutput
			}
		}

		fmt.Println(max)
	}
}
