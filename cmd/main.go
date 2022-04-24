package main

import (
	"fmt"

	"github.com/J-Obog/tcvm/pkg/assembler"
)

func main() {
	b := []byte("mov [r0] 5")
	l := assembler.Lexer{Input: b}

	for {
		tkn, _ := l.NextToken()
		if tkn == nil {
			break
		} 

		fmt.Println(tkn)
	}
}