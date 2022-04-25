package lexer

import "errors"

type Position struct {
	Column uint
	Line   uint
}

type Lexer struct {
	Input []byte
	Pos   Position
	index int
}

func IsAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func IsWhiteSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (lex *Lexer) curr() byte {
	if lex.index >= len(lex.Input) {
		return 0
	}
	return lex.Input[lex.index]
}

func (lex *Lexer) advance() {
	if lex.curr() == '\n' {
		lex.Pos.Column = 0
		lex.Pos.Line++
	} else {
		lex.Pos.Column++
	}
	lex.index++
}

func (lex *Lexer) lexNum() *Token {
	p := lex.Pos
	var buf string

	for IsDigit(lex.curr()){
		buf += string(lex.curr())
		lex.advance()
	}

	return &Token{Type: Number, Image: buf, Pos: p}
}

func (lex *Lexer) lexIdent() *Token {
	p := lex.Pos
	var buf string

	for IsAlpha(lex.curr()) || IsDigit(lex.curr()) || lex.curr() == '_' {
		buf += string(lex.curr())
		lex.advance()
	}

	return &Token{Type: Identifier, Image: buf, Pos: p}
}

func (lex *Lexer) NextToken() (*Token, error) {
	for IsWhiteSpace(lex.curr()) {
		lex.advance()
	}

	c := lex.curr()

	if c == 0 {
		return nil, nil
	}

	if IsDigit(c) {
		return lex.lexNum(), nil
	}

	if IsAlpha(c) || c == '_' {
		return lex.lexIdent(), nil
	} 

	if c == '[' || c == ']' {
		p := lex.Pos
		lex.advance()
		return &Token{Type: SpecialChar, Image: string(c), Pos: p}, nil
	}

	return nil, errors.New("Unrecognized token")
}
