package asm

import "github.com/J-Obog/tcvm/pkg/slf"

type Assembler struct {
	pgm *slf.Program
	par *Parser
}

func (a *Assembler) AssembleProgram() *slf.Program {
	stmt := a.par.NextStatement()

	for stmt != nil {
		switch stmt.(type) {
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

func (a *Assembler) handleData(data Statement) {
	//hadle any data we come across
}

func (a *Assembler) handleInstruction(op Statement) {
	//handle any instruction we come across
}

func (a *Assembler) handleLabel(label Statement) {
	//handle any label we come accross
}