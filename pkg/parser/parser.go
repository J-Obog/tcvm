package parser

import "github.com/J-Obog/tcvm/pkg/lexer"

type Intsruction struct {
	Type  uint8
	Size  uint8
	Value []byte
}

const ( // node type mapping
	instruction = iota //
	register
	effective
	label
	immediate
	data
)

type Parser struct {
	lex *lexer.Lexer //Lexer struct reference
}

func (parser *Parser) Parse() {

}