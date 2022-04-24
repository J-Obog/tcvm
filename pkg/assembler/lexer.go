package assembler

import (
	"bytes"
	"errors"
	"io"
	"unicode"
)

type Position struct {
	Column uint
	Line   uint
}

type Token struct {
	Type  int
	Image string
	Pos   Position
}

const ( //token type mapping
	Identifier = iota
	Number
	SpecialChar
)

type Lexer struct {
	Input string
	scanner *bytes.Reader
	Pos   Position
}

func (lex *Lexer) lexNum() *Token {
	var buf string
	p := lex.Pos
	for {
		r, _, err := lex.scanner.ReadRune()
		if err == io.EOF || !unicode.IsDigit(r) {
			return &Token{Type: Number, Image: buf, Pos: p}
		}
		buf += string(r)
	}
}

func (lex *Lexer) lexIdent() *Token {
	var buf string
	p := lex.Pos
	for {
		r, _, err := lex.scanner.ReadRune()
		if err == io.EOF || !(unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_') {
			return &Token{Type: Identifier, Image: buf, Pos: p}
		}
		buf += string(r)
	}
}

func (lex *Lexer) NextToken() (*Token, error) {
	for {
		r, _, err := lex.scanner.ReadRune()
		if err == io.EOF {
			return nil, nil
		} else if !unicode.IsSpace(r) {
			continue
		} else if unicode.IsDigit(r) {
			return lex.lexNum(), nil
		} else if unicode.IsLetter(r) || r == '_' {
			return lex.lexIdent(), nil
		} else if r == '[' {
			p := lex.Pos
			return &Token{Type: SpecialChar, Image: string(r), Pos: p}, nil
		} else {
			return nil, errors.New("Invalid Token")
		}
	}
}