package parser

import (
	"github.com/J-Obog/tcvm/pkg/lexer"
)

type Parser struct {
	tokens []*lexer.Token //list of tokens
	ct *lexer.Token //current token
	pos int //parser position
}

func New(lex *lexer.Lexer) *Parser {
	tkn := lex.NextToken()
	var st *lexer.Token
	p := &Parser{}

	for tkn != nil {
		p.tokens = append(p.tokens, tkn)
		tkn = lex.NextToken()
	}

	if len(p.tokens) > 0 {
		st = p.tokens[0]
	}

	p.ct = st
	return p
}

func (p *Parser) advance() {
	p.pos++

	if p.pos >= len(p.tokens) {
		p.ct = nil
	} else {
		p.ct = p.tokens[p.pos]
	}
}

func (p *Parser) parseLabel() Statement {
	p.advance()

	if p.ct == nil || p.ct.Type != lexer.Identifier {
		panic("Error parsing label")
	}

	lbl := &Label{Name: p.ct.Image}
	p.advance()


	return lbl
}

func (p *Parser) NextStatement() Statement {
	if p.ct == nil {
		return nil
	}

	if p.ct.Type == lexer.Identifier {
		txt := p.ct.Image

		if txt == "label" {
			return p.parseLabel()
		}
	}

	panic("Invalid statement")
}