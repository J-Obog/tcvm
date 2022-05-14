package asm

import "bytes"

type Lexer struct {
	buf *bytes.Buffer
}

type Token struct {
	Type  uint8
	Image string
}

const ( //token type mapping
	TKN_IDENTIFIER  uint8 = 0
	TKN_NUMBER      uint8 = 1
	TKN_REGISTER    uint8 = 2
	TKN_INSTRUCTION uint8 = 3
	TKN_LABEL       uint8 = 4
	TKN_ALLOCTYPE   uint8 = 5
	TKN_DATA        uint8 = 6
)

func (l *Lexer) NextToken() *Token {
	curr, _:= l.buf.ReadByte()

	if curr == 0 {
		return nil
	}

	for isWhiteSpace(curr) {
		curr, _ = l.buf.ReadByte()
	}

	if isDigit(curr) {
		return l.lexNum(curr)
	}

	if isAlpha(curr) || curr == '_' {
		return l.lexIdent(curr)
	}

	panic("Unrecognized token")
}

func (l *Lexer) lexNum(numByte byte) *Token {
	b := []byte{numByte}
	curr, _ := l.buf.ReadByte()

	for isDigit(curr) {
		b = append(b, curr)
		curr, _ = l.buf.ReadByte()
	}

	img := string(b)  
	l.buf.UnreadByte()

	return &Token{Type: TKN_NUMBER, Image: img}
}

func (l *Lexer) lexIdent(alphaByte byte) *Token {
	b := []byte{alphaByte}
	curr, _ := l.buf.ReadByte()

	for isAlpha(curr) || isDigit(curr) || curr == '_' {
		b = append(b, curr)
		curr, _ = l.buf.ReadByte()
	}

	img := string(b)  
	l.buf.UnreadByte()

	if img == "label" {
		return &Token{Type: TKN_LABEL, Image: img}
	}

	if img == "data" {
		return &Token{Type: TKN_DATA, Image: img}
	}

	if _, ok := REGISTER_TBL[img]; ok {
		return &Token{Type: TKN_REGISTER, Image: img}
	}

	if _, ok := INSTRUCTION_TBL[img]; ok {
		return &Token{Type: TKN_INSTRUCTION, Image: img}
	}

	if _, ok := ALLOCTYPE_TBL[img]; ok {
		return &Token{Type: TKN_ALLOCTYPE, Image: img}
	}

	return &Token{Type: TKN_IDENTIFIER, Image: img}
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}