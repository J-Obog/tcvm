package lexer

type Lexer struct {
	input []byte
	cb    byte //current byte
	pos   int  //lex position
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

	return &Token{Type: Number, Image: buf}
}

func (l *Lexer) lexIdent() *Token {
	var buf string

	for IsAlpha(l.cb) || IsDigit(l.cb) || l.cb == '_' {
		buf += string(l.cb)
		l.advance()
	}

	return &Token{Type: Identifier, Image: buf}
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

	if c == '[' || c == ']' {
		l.advance()
		return &Token{Type: SpecialChar, Image: string(c)}
	}

	panic("Unrecognized token")
}