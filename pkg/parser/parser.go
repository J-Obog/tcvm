package parser

import (
	"strconv"

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


func (p *Parser) parseData() Statement {
	p.advance()

	if p.ct == nil {
		panic("Unexpected EOF")
	}

	dsz, ok := dataTypes[p.ct.Image]

	if !ok {
		panic("Invalid datatype used in data definition")
	}

	p.advance()
	
	if p.ct == nil {
		panic("Unexpected EOF")
	}

	if p.ct.Type != lexer.Number {
		panic("Data must be of type num literal")
	}

    val, err := strconv.ParseUint(p.ct.Image, 10, 32)
    if err != nil {
        panic(err)
    }

	p.advance()

	return &Data{Size: dsz, Value: uint32(val)}
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

func (p *Parser) parseZeroOps() Statement {
	op := p.ct.Image
	p.advance()
	return &Instruction{Opcode: op}
} 

func (p *Parser) parseOneOp() Statement {
	return nil
} 

func (p *Parser) parseTwoOps() Statement {
	return nil
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

		if txt == "data" {
			return p.parseData()
		}

		if op, ok := opcodes[txt]; ok { //parsing an instruction
			a := op.Arity
			
			if a == 0 {
				return p.parseZeroOps()
			} else if a == 1 {
				return p.parseOneOp()
			} else {
				return p.parseTwoOps()
			}
		}
	}

	panic("Invalid statement")
}