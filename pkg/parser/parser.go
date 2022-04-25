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
	tokens []*lexer.Token //list of tokens
	ct *lexer.Token //current token
	pos int //parser position
}


func (parser *Parser) Parse() {

}