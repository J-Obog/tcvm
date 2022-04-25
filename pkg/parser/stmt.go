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
	Size  uint8
	Value uint32
}

func (dat *Data) TotalSize() uint8 {
	return dat.Size * 8
}

func (dat *Data) String() string {
	return fmt.Sprintf("[DATA %dB %d]", dat.Size, dat.Value)
}

const ( //operand modes
	ERegister = iota
	Register 
	Memory
	Immediate
)

type Operand struct {
	Size  uint8 //in bytes
	Mode  uint8
	Value string
}

type Instruction struct {
	Statement
	Opcode   string
	Operands []Operand
}

func (op *Instruction) String() string {
	return fmt.Sprintf("[INSTRUCTION %s %v]", op.Opcode, op.Operands)
}

func (op *Instruction) TotalSize() uint8 {
	sum := uint8(0)
	for _, opr := range op.Operands {
		sum += opr.Size
	}
	return sum * 8
}