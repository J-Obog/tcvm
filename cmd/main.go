package main

import (
	"fmt"

	"github.com/J-Obog/tcvm/pkg/lexer"
	"github.com/J-Obog/tcvm/pkg/parser"
)

func main() {
	l := lexer.New([]byte("nop mov8 [r5] 6 mov16 [r2] myloc"))
	p := parser.New(l)

	s := p.NextStatement()

	for s != nil {
		fmt.Println(s)
		s = p.NextStatement()
	}
}