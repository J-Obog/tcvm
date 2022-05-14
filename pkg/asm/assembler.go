package asm

import "github.com/J-Obog/tcvm/pkg/slf"

type Assembler struct {
	slf.Program
	code []Statement
}

func NewAssembler(p *Parser) *Assembler {
	a := &Assembler{}
	stmt := p.NextStatement()

	for stmt != nil {
		a.code = append(a.code, stmt)
		stmt = p.NextStatement()
	}

	return a
}