package lexer

type Token struct {
	Type  uint8
	Image string
}

const ( //token type mapping
	IDENTIFIER  uint8 = 0
	NUMBER      uint8 = 1
	REGISTER    uint8 = 2
	INSTRUCTION uint8 = 3
	LABEL       uint8 = 4
	ALLOCTYPE   uint8 = 5
	DATA        uint8 = 6
)