package asm

import "github.com/J-Obog/tcvm/pkg/slf"

type Assembler struct {
	pgm *slf.Program
	par *Parser
}

func (a *Assembler) AssembleProgram() *slf.Program {
	stmt := a.par.NextStatement()

	for stmt != nil {
		switch stmt := stmt.(type) {
		case *Label:
			a.handleLabel(stmt)
		case *Data:
			a.handleData(stmt) 
		case *Instruction:
			a.handleInstruction(stmt)
		}

		stmt = a.par.NextStatement()
	}

	return a.pgm
}

func (a *Assembler) handleData(data *Data) {
	//hadle any data we come across
}

func (a *Assembler) handleInstruction(op *Instruction) {
	//handle any instruction we come across
}

func (a *Assembler) handleLabel(label *Label) {
	lbl := label.Name
	sym := a.pgm.SymTab[lbl]  

	if sym == nil {
		a.pgm.StrTab = append(a.pgm.StrTab, lbl)
		newSym := &slf.Symbol{}
		newSym.StrTabIndex = uint32(len(a.pgm.StrTab) - 1)
		newSym.Offset = uint32(len(a.pgm.CodeSeg))
		a.pgm.SymTab[lbl] = sym
	} else {
		if ((sym.Flags >> slf.S_EXTERN) & 0x1) == 1 {
				sym.Offset = uint32(len(a.pgm.CodeSeg))
		} else {
			panic("Redefinition of symbol")
		}
	}

}

