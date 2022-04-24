package main

import (
	"bytes"
	"fmt"

	"github.com/J-Obog/tcvm/pkg/assembler"
)

func main() {
	b := bytes.NewReader([]byte("[mov16 r0 12345]"))
	l := assembler.Lexer{Scanner: b}

	tkn, _ := l.NextToken()
	fmt.Println(tkn)
}