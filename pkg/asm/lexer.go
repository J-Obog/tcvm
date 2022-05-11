package asm

type Lexer struct {
	input []byte
	cb    byte //current byte
	pos   int  //lex position
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
	return &Lexer{input: input, cb: sb}
}

func (l *Lexer) advance() {
	l.pos++

	if l.pos >= len(l.input) {
		l.cb = 0
	} else {
		l.cb = l.input[l.pos]
	}
}

func (l *Lexer) lexNum() *Token {
	var buf string

	for IsDigit(l.cb) {
		buf += string(l.cb)
		l.advance()
	}

	return &Token{Type: TKN_NUMBER, Image: buf}
}

func (l *Lexer) lexIdent() *Token {
	var buf string

	for IsAlpha(l.cb) || IsDigit(l.cb) || l.cb == '_' {
		buf += string(l.cb)
		l.advance()
	}

	return &Token{Type: TKN_IDENTIFIER, Image: buf}
}

func (l *Lexer) NextToken() *Token {
	for IsWhiteSpace(l.cb) {
		l.advance()
	}

	c := l.cb

	if c == 0 {
		return nil
	}

	if IsDigit(c) {
		return l.lexNum()
	}

	if IsAlpha(c) || c == '_' {
		return l.lexIdent()
	}

	panic("Unrecognized token")
}
