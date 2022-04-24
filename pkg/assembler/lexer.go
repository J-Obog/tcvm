package assembler

import (
	"regexp"
)

type Position struct {
	Column uint
	Line   uint
}

type Token struct {
	Type  int
	Image string
	Pos   Position
}

const ( //token type mapping
	Identifier = iota
	Number
	SpecialChar
)

type Lexer struct {
	Input string
	index int
	Pos   Position
}

func NewLexer(input string) *Lexer {
	rxp, _ := regexp.Compile(`\;.*\n`)
	stripped := rxp.ReplaceAllString(input, " ")
	return &Lexer{Input: stripped}
}