package main

import (
	"fmt"

	"github.com/J-Obog/tcvm/pkg/lexer"
	"github.com/J-Obog/tcvm/pkg/parser"
)

func main() {
	l := lexer.New([]byte("data b 122"))
	p := parser.New(l)

	st := p.NextStatement()
	fmt.Println(st)
}