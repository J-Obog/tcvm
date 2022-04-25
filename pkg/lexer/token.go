package lexer

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