package main

import (
	"fmt"

	"github.com/J-Obog/tcvm/pkg/lexer"
	"github.com/J-Obog/tcvm/pkg/parser"
)

func main() {
	l := lexer.New([]byte("label myloc\n\tdata b 122\n\nlabel __start__ nop"))
	p := parser.New(l)

	s := p.NextStatement()

	for s != nil {
		fmt.Println(s)
		s = p.NextStatement()
	}
}