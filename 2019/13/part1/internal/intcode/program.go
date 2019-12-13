package intcode

import (
	"strconv"
)

type Memory map[int]int

type InstructionFunc func(i Instruction)

type Program struct {
	memory   Memory
	base     int
	position int

	ops map[InstructionType]InstructionFunc

	ioChan IOChan
}

func NewProgram(definition []string) (*Program, error) {
	memory := make(Memory, len(definition))
	for i := range definition {
		n, err := strconv.Atoi(definition[i])
		if err != nil {
			return nil, err
		}

		memory[i] = n
	}

	p := &Program{memory: memory}
	p.ops = map[InstructionType]InstructionFunc{
		SUM: p.doSum, MUL: p.doMul, IN: p.doInput, OUT: p.doOutput,
		JT: p.doJumpTrue, JF: p.doJumpFalse, LT: p.doLessThan,
		EQ: p.doEqual, BASE: p.doBase,
	}

	return p, nil
}

func (p *Program) Execute(ioChan IOChan) {
	p.ioChan = ioChan

	i := NewInstruction(p.position, p.base, p.memory)

	for !i.Type.IsHalt() {
		p.do(i)
		i = NewInstruction(p.position, p.base, p.memory)
	}

	p.ioChan.Halt()
}

func (p *Program) do(i Instruction) {
	p.ops[i.Type](i)
}

func (p *Program) doSum(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)
	res := i.MemPos(3)

	p.memory[res] = x + y

	p.position += 4
}

func (p *Program) doMul(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)
	res := i.MemPos(3)

	p.memory[res] = x * y

	p.position += 4
}

func (p *Program) doLessThan(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)
	res := i.MemPos(3)

	if x < y {
		p.memory[res] = 1
	} else {
		p.memory[res] = 0
	}

	p.position += 4
}

func (p *Program) doEqual(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)
	res := i.MemPos(3)

	if x == y {
		p.memory[res] = 1
	} else {
		p.memory[res] = 0
	}

	p.position += 4
}

func (p *Program) doJumpFalse(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)

	if x == 0 {
		p.position = y
	} else {
		p.position += 3
	}
}

func (p *Program) doJumpTrue(i Instruction) {
	x := i.Param(1)
	y := i.Param(2)

	if x != 0 {
		p.position = y
	} else {
		p.position += 3
	}
}

func (p *Program) doBase(i Instruction) {
	p.position += 2
	p.base += i.Param(1)
}

func (p *Program) doInput(i Instruction) {
	input := p.ioChan.Receive()
	i.memory[i.MemPos(1)] = input
	p.position += 2
}

func (p *Program) doOutput(i Instruction) {
	p.ioChan.Send(i.Param(1))
	p.position += 2
}
