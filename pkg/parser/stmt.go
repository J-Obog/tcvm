package parser

import "fmt"

type Statement interface {
	String() string
	TotalSize() uint32 //in bytes
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

func (dat *Data) TotalSize() uint32 {
	switch dat.Specifier {
	case Byte:
		return 1
	case Word:
		return 2
	case DWord:
		return 4
	case Space:
		return dat.Value
	default:
		return 0xFFFFFFFF
	}
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
	return fmt.Sprintf("[INSTRUCTION %d %v]", op.Opcode, op.Operands)
}

func (op *Instruction) TotalSize() uint32 {
	arity := len(op.Operands)
	
	switch arity {
	case 0:
		return 1

	case 1:
		opr := op.Operands[0]
		if opr.Source == Register {
			return 2
		} else {
			return 5
		}

	case 2:
		opr1 := op.Operands[0]
		opr2 := op.Operands[1]

		if opr1.Source == Register && opr2.Source == Register {
			return 2
		} else {
			return 7
		}

	default:
		return 0xFFFFFFFF
	}
}