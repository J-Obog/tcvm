package assembler

type Position struct {
	Column uint
	Line   uint
}

type Token struct {
	Type  int
	Image string
	Pos   Position
}

type Lexer struct {
	Input []byte
	index uint
	Pos   Position
}