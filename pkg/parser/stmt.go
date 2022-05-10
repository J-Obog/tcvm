package parser

import "fmt"

type Statement interface {
	String() string
	TotalSize() uint8 //in bytes
}

type Label struct {
	Statement
	Name string
}

func (lbl *Label) String() string {
	return fmt.Sprintf("[LABEL %s]", lbl.Name)
}

func (lbl *Label) TotalSize() uint8 {
	return 0
}

type Data struct {
	Statement
	Specifier uint8
	Value uint32
	VarName string
}

func (dat *Data) TotalSize() uint8 {
	return 0//dat.Size * 8
}

func (dat *Data) String() string {
	return fmt.Sprintf("[DATA %dB %d]", dat.Specifier, dat.Value)
}

type Operand struct {
	Source uint8 //REG | MEM | IMM	
	Value string
}

type Instruction struct {
	Statement
	Opcode   uint8
	Operands []Operand
}

func (op *Instruction) String() string {
	return fmt.Sprintf("[INSTRUCTION %s %v]", op.Opcode, op.Operands)
}

func (op *Instruction) TotalSize() uint8 {
	return 0
	/*sum := uint8(0)
	for _, opr := range op.Operands {
		sum += opr.Size
	}
	return sum * 8*/
}