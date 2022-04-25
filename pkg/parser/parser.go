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

	lbl := p.ct.Image
	p.advance()

	return &Label{Name: lbl}
}

func (p *Parser) getOperand() Operand {
	p.advance()

	if p.ct == nil {
		panic("Unexpected EOF")
	}
	
	if p.ct.Type == lexer.Number {
		return Operand{Size: 4, Mode: Immediate, Value: p.ct.Image} 
	}

	if p.ct.Type == lexer.Identifier {
		_, ok := registers[p.ct.Image]

		if ok {
			return Operand{Size: 1, Mode: Register, Value: p.ct.Image}
		}

		return Operand{Size: 4, Mode: Memory, Value: p.ct.Image}
	}

	if p.ct.Image == "[" {
		p.advance()
		
		if p.ct == nil {
			panic("Unexpected EOF")
		}

		r := p.ct.Image
		_, ok := registers[r]
		if !ok || r[0] != 'r' { //if register is invalid or trying to dereference a special purpose register
			panic("Invalid effective address")
		}

		p.advance()

		if p.ct == nil {
			panic("Unexpected EOF")
		}

		if p.ct.Image == "]" {
			return Operand{Size: 1, Mode: ERegister, Value: r}
		}
	}

	panic("Invalid source")
} 

func (p *Parser) parseInstruction() Statement {
	opc := p.ct.Image
	a := opcodes[opc].Arity 

	if a == 0 { // zero operand instruction
		return &Instruction{Opcode: opc}	
	} else if a == 1 {// one operand instruction
		src := p.getOperand() 
		if opc == "not" && src.Mode != Register {
			panic("Invalid source type")
		}	
		if opc == "push8" {
			src.Size = 1
		}
		if opc == "push16" {
			src.Size = 2
		}

		return &Instruction{Opcode: opc, Operands: []Operand{src}}
	} else { // two operand instruction
		return nil
	}
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

		_, ok := opcodes[txt] 
		if ok {
			return p.parseInstruction()
		}
	}

	panic("Invalid statement")
}