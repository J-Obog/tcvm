package lexer

type Token struct {
	Type  int
	Image string
}

const ( //token type mapping
	Identifier = iota
	Number
)