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

}

func (p *Parser) parseLabel() Statement {

}

func (p *Parser) parseInstruction() Statement {

}

func (p *Parser) NextStatement() Statement {
	if p.ct == nil {
		return nil
	}

	switch p.ct.Type {
	case TKN_LABEL:
		return nil

	case TKN_DATA:
		return nil

	case TKN_INSTRUCTION:
		return nil
	}

	panic("Invalid statement")
}