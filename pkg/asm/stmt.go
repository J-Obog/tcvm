package asm

import "fmt"

type Statement interface {
	String() string
}

type Label struct {
	Statement
	Name string
}

func (lbl *Label) String() string {
	return fmt.Sprintf("[LABEL %s]", lbl.Name)
}

type Data struct {
	Statement
	AllocType uint8
	Literal string
	LabelId string
}

func (dat *Data) String() string {
	return fmt.Sprintf("[DATA %s %d %s]", dat.LabelId, dat.AllocType, dat.Literal)
}

type Operand struct {
	OperandType uint8
	Literal string
}

type Instruction struct {
	Statement
	Opcode   uint8
	Operands []Operand
}

func (op *Instruction) String() string {
	return fmt.Sprintf("[INSTRUCTION %d %v]", op.Opcode, op.Operands)
}
