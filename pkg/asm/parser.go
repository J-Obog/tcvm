package asm

type Parser struct {
	Lex *Lexer
}

func (p *Parser) NextStatement() Statement {
	curr := p.Lex.NextToken()

	if curr == nil {
		return nil
	}

	switch curr.Type {
	case TKN_LABEL:
		return p.parseLabel()

	case TKN_DATA:
		return p.parseData()

	case TKN_INSTRUCTION:
		return p.parseInstruction(curr)
	}

	panic("Invalid statement")
}

func (p *Parser) parseData() Statement {
	data := &Data{}
	curr := p.Lex.NextToken()

	if curr == nil {
		panic("Unexpected EOF")
	}

	if curr.Type == TKN_IDENTIFIER { //optional data label
		data.LabelId = curr.Image
		curr = p.Lex.NextToken()
		if curr == nil {
			panic("Unexpected EOF")
		}
	}

	if curr.Type != TKN_ALLOCTYPE { //alloctype must follow data directive or data label
		panic("Invalid specifier used in data definition")
	} else {
		data.AllocType = ALLOCTYPE_TBL[curr.Image]
		curr = p.Lex.NextToken()
		if curr == nil {
			panic("Unexpected EOF")
		}
	}

	if curr.Type != TKN_NUMBER { //num literal must follow alloctype
		panic("Data value must be of type num literal")
	} else {
		data.Literal = curr.Image
	}

	return data
}

func (p *Parser) parseLabel() Statement {
	curr := p.Lex.NextToken()

	if curr == nil {
		panic("Unexpected EOF")
	}

	if curr.Type != TKN_IDENTIFIER { //label name must be a valid identifier
		panic("Error parsing label")
	}

	return &Label{Name: curr.Image}
}

func (p *Parser) isOperandType(oprType uint8) bool {
	return oprType == (TKN_IDENTIFIER) || (oprType == TKN_REGISTER) || (oprType == TKN_NUMBER)
}

func (p *Parser) parseInstruction(opTkn *Token) Statement {
	op := &Instruction{Opcode: INSTRUCTION_TBL[opTkn.Image]}

	curr := p.Lex.NextToken()

	for curr != nil && p.isOperandType(curr.Type) { //consume all possible operands
		opr := Operand{OperandType: curr.Type, Literal: curr.Image}
		op.Operands = append(op.Operands, opr)
		curr = p.Lex.NextToken()
	}

	return op
}