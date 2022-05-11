package asm

type Parser struct {
	tokens []*Token //list of tokens
	ct     *Token   //current token
	pos    int      //parser position
}

func NewParser(lex *Lexer) *Parser {
	tkn := lex.NextToken()
	var st *Token
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
	data := &Data{}

	if p.ct == nil {
		panic("Unexpected EOF")
	}

	if p.ct.Type == TKN_IDENTIFIER { //optional data label
		data.LabelId = p.ct.Image
		p.advance()
		if p.ct == nil {
			panic("Unexpected EOF")
		}
	}

	if p.ct.Type != TKN_ALLOCTYPE { //alloctype must follow data directive or data label
		panic("Invalid specifier used in data definition")
	} else {
		data.AllocType = ALLOCTYPE_TBL[p.ct.Image]
		p.advance()
		if p.ct == nil {
			panic("Unexpected EOF")
		}
	}

	if p.ct.Type != TKN_NUMBER { //num literal must follow alloctype
		panic("Data value must be of type num literal")
	} else {
		data.Literal = p.ct.Image
		p.advance()
	}

	return data
}

func (p *Parser) parseLabel() Statement {
	p.advance()

	if p.ct == nil {
		panic("Unexpected EOF")
	}

	if p.ct.Type != TKN_IDENTIFIER { //label name must be a valid identifier
		panic("Error parsing label")
	}

	nm := p.ct.Image
	p.advance()
	return &Label{Name: nm}
}

func (p *Parser) isOperandType(oprType uint8) bool {
	return oprType == (TKN_IDENTIFIER) || (oprType == TKN_REGISTER) || (oprType == TKN_NUMBER)
}

func (p *Parser) parseInstruction() Statement {
	op := &Instruction{Opcode: INSTRUCTION_TBL[p.ct.Image]}
	p.advance()

	for p.ct != nil && p.isOperandType(p.ct.Type) { //consume all possible operands
		opr := Operand{OperandType: p.ct.Type, Literal: p.ct.Image}
		op.Operands = append(op.Operands, opr)
		p.advance()
	}

	return op
}

func (p *Parser) NextStatement() Statement {
	if p.ct == nil {
		return nil
	}

	switch p.ct.Type {
	case TKN_LABEL:
		return p.parseLabel()

	case TKN_DATA:
		return p.parseData()

	case TKN_INSTRUCTION:
		return p.parseInstruction()
	}

	panic("Invalid statement")
}