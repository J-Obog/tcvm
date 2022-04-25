package assembler

import (
	"github.com/J-Obog/tcvm/pkg/parser"
)

type Assembler struct {
	size    uint32
	symtab  map[string]uint32
	program []parser.Statement
}

func New(p *parser.Parser) *Assembler {
	return nil
}