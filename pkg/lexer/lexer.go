package lexer

import "errors"

type Lexer struct {
	Input []byte
	cb byte //current byte
	pos int //lex position
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

func New(input []byte) *Lexer {
	var sb byte

	if len(input) > 0 {
		sb = input[0]	
	} 
	return &Lexer{Input: input, cb: sb}
}


func (lex *Lexer) advance() {
	lex.pos++
	
	if lex.pos >= len(lex.Input) {
		lex.cb = 0
	} else {
		lex.cb = lex.Input[lex.pos]
	}
}

func (lex *Lexer) lexNum() *Token {
	var buf string

	for IsDigit(lex.cb){
		buf += string(lex.cb)
		lex.advance()
	}

	return &Token{Type: Number, Image: buf}
}

func (lex *Lexer) lexIdent() *Token {
	var buf string

	for IsAlpha(lex.cb) || IsDigit(lex.cb) || lex.cb == '_' {
		buf += string(lex.cb)
		lex.advance()
	}

	return &Token{Type: Identifier, Image: buf}
}

func (lex *Lexer) NextToken() (*Token, error) {
	for IsWhiteSpace(lex.cb) {
		lex.advance()
	}

	c := lex.cb

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
		lex.advance()
		return &Token{Type: SpecialChar, Image: string(c)}, nil
	}

	return nil, errors.New("Unrecognized token")
}
