package intcode

const (
	// ops
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
	position  = 0
	immediate = 1
	relative  = 2
)

type InstructionType int

func (it InstructionType) IsHalt() bool {
	return it == HALT
}

type Instruction struct {
	Type InstructionType

	currPos  int
	currBase int
	memory   Memory

	modes []int
}

func NewInstruction(currPos, currBase int, memory Memory) Instruction {
	op := memory[currPos]

	return Instruction{
		Type: InstructionType(op % 100),

		currPos:  currPos,
		currBase: currBase,
		memory:   memory,

		modes: []int{(op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10},
	}
}

func (i Instruction) Arg(param int) int {
	return i.memory[i.currPos+param]
}

func (i Instruction) Param(param int) int {
	if i.modes[param-1] == immediate {
		return i.Arg(param)
	}

	if i.modes[param-1] == relative {
		return i.memory[i.currBase+i.Arg(param)]
	}

	return i.memory[i.Arg(param)]
}

func (i Instruction) MemPos(param int) int {
	if i.modes[param-1] == position {
		return i.Arg(param)
	} else {
		return i.currBase + i.Arg(param)
	}
}
