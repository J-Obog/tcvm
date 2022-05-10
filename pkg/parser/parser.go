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
		return Operand{Source: Immediate, Value: p.ct.Image} 
	}

	if p.ct.Type == lexer.Identifier {
		_, ok := registers[p.ct.Image]

		if ok {
			return Operand{Source: Register, Value: p.ct.Image}
		}

		return Operand{Source: Memory, Value: p.ct.Image}
	}

	panic("Invalid operand type")
} 

func (p *Parser) parseInstruction(opcode uint8) Statement {
	primaryOp := (opcode >> 5) & 0x7

	switch primaryOp {
	case Nop, SysCall:
		p.advance()
		return &Instruction{Opcode: opcode}

	case DTransfer, Alu:
		op1 := p.getOperand()
		op2 := p.getOperand()

		if op1.Source != Register {
			panic("Invalid combination of opcode and operands")
		}

		p.advance()
		return &Instruction{Opcode: opcode, Operands: []Operand{op1, op2}}

	case Jump:
		op := p.getOperand()
		p.advance()
		return &Instruction{Opcode: opcode, Operands: []Operand{op}}

	default:
		panic("Invalid opcode encoding") 
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

		opc, ok := opcodeMap[txt]
		if ok {
			return p.parseInstruction(opc)
		}
	}

	panic("Invalid statement")
}